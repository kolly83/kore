import * as React from 'react'
import PropTypes from 'prop-types'
import { patterns } from '../../../utils/validation'
import { Button, Form, Icon, Input, Alert, Select, Typography } from 'antd'
const { Paragraph } = Typography
import KoreApi from '../../../kore-api'
import DataField from '../../../components/utils/DataField'
import UsePlanForm from '../../plans/UsePlanForm'
import V1ServiceCredentials from '../../../kore-api/model/V1ServiceCredentials'
import V1ServiceCredentialsSpec from '../../../kore-api/model/V1ServiceCredentialsSpec'
import { NewV1ObjectMeta, NewV1Ownership } from '../../../utils/model'
import FormErrorMessage from '../../forms/FormErrorMessage'

class ServiceCredentialForm extends React.Component {
  static propTypes = {
    form: PropTypes.any.isRequired,
    creationSource: PropTypes.oneOf(['namespace', 'service']).isRequired,
    team: PropTypes.object.isRequired,
    clusters: PropTypes.array,
    namespaceClaims: PropTypes.array,
    services: PropTypes.array,
    handleSubmit: PropTypes.func.isRequired,
    handleCancel: PropTypes.func.isRequired
  }

  constructor(props) {
    super(props)
    this.state = {
      clusters: props.clusters,
      services: props.services,
      namespaceClaims: props.namespaceClaims,
      servicePlan: null,
      submitting: false,
      formErrorMessage: false,
      dataLoading: true,
      validationErrors: null,
      config: null
    }
  }

  componentDidMountComplete = null
  componentDidMount() {

    // Assign the promise chain to a variable so tests can wait for it to complete.
    this.componentDidMountComplete = Promise.resolve().then(async () => {
      const team = this.props.team.metadata.name
      const api = await KoreApi.client()
      switch (this.props.creationSource) {
      case 'namespace': {
        let [services, serviceKinds] = await Promise.all([
          api.ListServices(team),
          api.ListServiceKinds()
        ])
        services = services.items.filter(s => !s.spec.cluster || !s.spec.cluster.name)
        services = services.map(s => ({
          ...s,
          serviceKind: serviceKinds.items.find(sk => sk.metadata.name === s.spec.kind)
        }))
        this.setState({ services, dataLoading: false })
        this.props.form.validateFields()
        break
      }
      case 'service': {
        let [clusters, namespaceClaims, serviceKinds] = await Promise.all([
          api.ListClusters(team),
          api.ListNamespaces(team),
          api.ListServiceKinds()
        ])
        clusters = clusters.items.filter(cluster => namespaceClaims.items.filter(ns => ns.spec.cluster.name === cluster.metadata.name).length > 0)
        if (clusters.length === 1) {
          namespaceClaims = namespaceClaims.items.filter(nc => nc.spec.cluster.name === clusters[0].metadata.name)
        } else {
          namespaceClaims = namespaceClaims.items
        }
        this.setState({ clusters, namespaceClaims, serviceKinds, dataLoading: false })
        this.props.form.validateFields()
        break
      }
      }
    })
  }

  disableButton = fieldsError => {
    if (this.state.submitting) {
      return true
    }
    return Object.keys(fieldsError).some(field => fieldsError[field])
  }

  cancel = () => {
    this.props.form.resetFields()
    this.props.handleCancel()
  }

  onClusterChange = () => {
    this.props.form.resetFields(['namespace'])
    this.props.form.validateFields()
  }

  onServiceChange = async (serviceName) => {
    const service = this.state.services.find(s => s.metadata.name === serviceName)

    const api = await KoreApi.client()
    const servicePlan = await api.GetServicePlan(service.spec.plan)
    this.setState({
      servicePlan: servicePlan,
      config: null,
      validationErrors: null
    })
  }

  handleConfigurationUpdate = (config) => this.setState({ config })

  handleSubmit = e => {
    e.preventDefault()

    const { clusters, services, namespaceClaims, config } = this.state

    this.setState({
      submitting: true,
      formErrorMessage: false
    })

    return this.props.form.validateFields(async (err, values) => {
      if (err) {
        return this.setState({ submitting: false })
      }

      const service = services.length === 1 ? services[0] : services.find(s => s.metadata.name === values.service)
      const cluster = clusters.length === 1 ? clusters[0] : clusters.find(c => c.metadata.name === values.cluster)
      const namespaceClaim = namespaceClaims.length === 1 ? namespaceClaims[0] : namespaceClaims.find(n => n.spec.name === values.namespace)
      const name = `${cluster.metadata.name}-${namespaceClaim.spec.name}-${values.secretName}`

      try {
        const existing = await (await KoreApi.client()).GetServiceCredentials(this.props.team.metadata.name, name)
        if (existing) {
          return this.setState({
            submitting: false,
            formErrorMessage: `This namespace already contains a service access with secret name "${values.secretName}"`,
            validationErrors: null
          })
        }
      } catch (err) {
        // TODO: we should differentiate between 404 and other errors
      }

      const serviceCredential = new V1ServiceCredentials()
      serviceCredential.setApiVersion('servicecredentials.services.kore.appvia.io/v1')
      serviceCredential.setKind('ServiceCredentials')
      serviceCredential.setMetadata(NewV1ObjectMeta(name, this.props.team.metadata.name))

      const serviceCredentialSpec = new V1ServiceCredentialsSpec()
      serviceCredentialSpec.setKind(service.spec.kind)
      serviceCredentialSpec.setService(NewV1Ownership({
        group: service.apiVersion.split('/')[0],
        version: service.apiVersion.split('/')[1],
        kind: service.kind,
        name: service.metadata.name,
        namespace: this.props.team.metadata.name
      }))
      serviceCredentialSpec.setCluster(NewV1Ownership({
        group: cluster.apiVersion.split('/')[0],
        version: cluster.apiVersion.split('/')[1],
        kind: cluster.kind,
        name: cluster.metadata.name,
        namespace: this.props.team.metadata.name
      }))
      serviceCredentialSpec.setClusterNamespace(namespaceClaim.spec.name)
      serviceCredentialSpec.setSecretName(values.secretName)
      serviceCredentialSpec.setConfiguration(config)

      serviceCredential.setSpec(serviceCredentialSpec)

      try {
        const result = await (await KoreApi.client()).UpdateServiceCredentials(this.props.team.metadata.name, name, serviceCredential)
        this.props.form.resetFields()
        this.setState({ submitting: false })
        await this.props.handleSubmit(result)
      } catch (err) {
        this.setState({
          submitting: false,
          formErrorMessage: (err.fieldErrors && err.message) ? err.message : 'An error occurred creating service access, please try again',
          validationErrors: err.fieldErrors // This will be undefined on non-validation errors, which is fine.
        })
      }
    })
  }

  createFromNamespace = () => {
    const { getFieldDecorator } = this.props.form
    const { services, clusters, namespaceClaims } = this.state

    const namespace = namespaceClaims[0].spec.name
    const cluster = clusters[0].metadata.name

    return (
      <>
        <Alert
          message="Choose the service and secret name"
          description={
            <>
              <Paragraph>A secret with the chosen name will be populated in the namespace. The secret will contain credentials that can be used to access the service from within the namespace.</Paragraph>
              <DataField label="Cluster" value={cluster}/>
              <DataField label="Namespace" value={namespace}/>
            </>
          }
          type="info"
          showIcon
          style={{ marginBottom: '20px' }}
        />
        <Form.Item label="Service" validateStatus={this.fieldError('service') ? 'error' : ''} help={this.fieldError('service') || ''}>
          {getFieldDecorator('service', {
            rules: [{ required: true, message: 'Please select the service!' }],
            initialValue: services.length === 1 ? services[0].metadata.name : undefined
          })(
            <Select placeholder="Service" onChange={this.onServiceChange}>
              {services.map(s => (
                <Select.Option key={s.metadata.name} value={s.metadata.name}>{s.serviceKind && s.serviceKind.spec.displayName} - {s.metadata.name}</Select.Option>
              ))}
            </Select>
          )}
        </Form.Item>
        {this.secretNameField()}
      </>
    )
  }

  secretNameField = () => (
    <Form.Item label="Secret Name" validateStatus={this.fieldError('secretName') ? 'error' : ''} help={this.fieldError('secretName') || ''}>
      {this.props.form.getFieldDecorator('secretName', {
        rules: [
          { required: true, message: 'Please enter the secret name!' },
          { ...patterns.uriCompatible63CharMax }
        ]
      })(
        <Input placeholder="The name of the Kubernetes Secret"/>
      )}
    </Form.Item>
  )

  createFromService = () => {
    const { getFieldDecorator } = this.props.form
    const { clusters, namespaceClaims } = this.state

    return (
      <>
        <Alert
          message="Choose the cluster and namespace"
          description={
            <Paragraph>A secret with the chosen name will be populated in the namespace. The secret will contain credentials that can be used to access the service from within the namespace.</Paragraph>
          }
          type="info"
          showIcon
          style={{ marginBottom: '20px' }}
        />

        <Form.Item label="Cluster" validateStatus={this.fieldError('cluster') ? 'error' : ''} help={this.fieldError('cluster') || ''}>
          {getFieldDecorator('cluster', {
            rules: [{ required: true, message: 'Please select the cluster!' }],
            initialValue: clusters.length === 1 ? clusters[0].metadata.name : undefined
          })(
            <Select placeholder="Cluster" onChange={this.onClusterChange}>
              {clusters.map(c => (
                <Select.Option key={c.metadata.name} value={c.metadata.name}>{c.metadata.name}</Select.Option>
              ))}
            </Select>
          )}
        </Form.Item>
        <Form.Item label="Namespace" validateStatus={this.fieldError('namespace') ? 'error' : ''} help={this.fieldError('namespace') || ''}>
          {getFieldDecorator('namespace', {
            rules: [{ required: true, message: 'Please select the namespace!' }],
            initialValue: namespaceClaims.length === 1 ? namespaceClaims[0].spec.name : undefined
          })(
            <Select placeholder="Namespace">
              {namespaceClaims.filter(nc => nc.spec.cluster.name === this.props.form.getFieldValue('cluster')).map(n => (
                <Select.Option key={n.spec.name} value={n.spec.name}>{n.spec.name}</Select.Option>
              ))}
            </Select>
          )}
        </Form.Item>
        {this.secretNameField()}
      </>
    )
  }

  // Only show error after a field is touched.
  fieldError = fieldKey => this.props.form.isFieldTouched(fieldKey) && this.props.form.getFieldError(fieldKey)

  render() {
    if (this.state.dataLoading) {
      return <Icon type="loading"/>
    }

    const { getFieldsError } = this.props.form
    const { submitting, formErrorMessage, servicePlan, validationErrors } = this.state
    const formConfig = {
      layout: 'horizontal',
      labelAlign: 'left',
      hideRequiredMark: true,
      labelCol: {
        sm: { span: 24 },
        md: { span: 6 },
        lg: { span: 4 }
      },
      wrapperCol: {
        span: 12
      }
    }

    return (
      <Form {...formConfig} onSubmit={this.handleSubmit} style={{ marginBottom: '30px' }}>
        <FormErrorMessage message={formErrorMessage}/>

        {this.props.creationSource === 'namespace' && this.createFromNamespace() }
        {this.props.creationSource === 'service' && this.createFromService() }

        {servicePlan && (
          <UsePlanForm
            team={this.props.team}
            resourceType="servicecredential"
            kind={servicePlan.spec.kind}
            plan={servicePlan.metadata.name}
            mode="create"
            validationErrors={validationErrors}
            onPlanChange={this.handleConfigurationUpdate}
          />
        )}

        <Form.Item>
          <Button type="primary" htmlType="submit" loading={submitting} disabled={this.disableButton(getFieldsError())}>Save</Button>
          <Button type="link" onClick={this.cancel}>Cancel</Button>
        </Form.Item>
      </Form>
    )
  }
}

const WrappedServiceCredentialForm = Form.create({ name: 'service_credential' })(ServiceCredentialForm)

export default WrappedServiceCredentialForm
