import * as React from 'react'
import PropTypes from 'prop-types'
import { Alert, Button, Card, Form, Icon, message, Select } from 'antd'

import ServiceOptionsForm from '../../../services/ServiceOptionsForm'
import FormErrorMessage from '../../../forms/FormErrorMessage'
import KoreApi from '../../../../kore-api'
import V1ServiceSpec from '../../../../kore-api/model/V1ServiceSpec'
import V1Service from '../../../../kore-api/model/V1Service'
import { NewV1ObjectMeta, NewV1Ownership } from '../../../../utils/model'
import { getKoreLabel } from '../../../../utils/crd-helpers'

class ClusterApplicationServiceForm extends React.Component {
  static propTypes = {
    form: PropTypes.any.isRequired,
    team: PropTypes.object.isRequired,
    cluster: PropTypes.object.isRequired,
    teamServices: PropTypes.array.isRequired,
    handleSubmit: PropTypes.func.isRequired,
    handleCancel: PropTypes.func.isRequired
  }

  static initialState = {
    submitButtonText: 'Save',
    submitting: false,
    formErrorMessage: false,
    selectedServiceKind: false,
    dataLoading: true,
    servicePlanOverride: null,
    validationErrors: null
  }

  state = { ...ClusterApplicationServiceForm.initialState }

  async fetchComponentData() {
    const api = await KoreApi.client()
    const [ serviceKinds, servicePlans, namespaceClaims ] = await Promise.all([
      api.ListServiceKinds(),
      api.ListServicePlans(),
      api.ListNamespaces(this.props.team.metadata.name)
    ])
    serviceKinds.items = serviceKinds.items.filter(sk => getKoreLabel(sk, 'platform') === 'Kubernetes' && sk.spec.enabled)
    namespaceClaims.items = namespaceClaims.items.filter(nc => nc.spec.cluster.name === this.props.cluster.metadata.name)
    return { serviceKinds, servicePlans, namespaceClaims }
  }

  componentDidMountComplete = null
  componentDidMount() {
    // Assign the promise chain to a variable so tests can wait for it to complete.
    this.componentDidMountComplete = Promise.resolve().then(async () => {
      const data = await this.fetchComponentData()
      this.setState({ ...data, dataLoading: false })
    })
  }

  generateServiceResource = (values) => {
    const cluster = this.props.cluster
    const selectedServicePlan = this.state.servicePlans.items.find(p => p.metadata.name === values.servicePlan)

    const serviceResource = new V1Service()
    serviceResource.setApiVersion('services.compute.kore.appvia.io/v1')
    serviceResource.setKind('Service')

    serviceResource.setMetadata(NewV1ObjectMeta(values.serviceName, this.props.team.metadata.name))

    const serviceSpec = new V1ServiceSpec()
    serviceSpec.setKind(selectedServicePlan.spec.kind)
    serviceSpec.setPlan(selectedServicePlan.metadata.name)
    serviceSpec.setClusterNamespace(values.namespace)

    serviceSpec.setCluster(NewV1Ownership({
      group: cluster.apiVersion.split('/')[0],
      version: cluster.apiVersion.split('/')[1],
      kind: cluster.kind,
      name: cluster.metadata.name,
      namespace: this.props.team.metadata.name
    }))
    if (this.state.servicePlanOverride) {
      serviceSpec.setConfiguration(this.state.servicePlanOverride)
    } else {
      serviceSpec.setConfiguration({ ...selectedServicePlan.spec.configuration })
    }

    serviceResource.setSpec(serviceSpec)
    return serviceResource
  }

  validatedFormsFields = (callback) => {
    this.props.form.validateFields((serviceErr, serviceValues) => {
      this.serviceOptionsForm.props.form.validateFields((optionsErr, optionsValues) => {
        const err = serviceErr || optionsErr ? { ...serviceErr, ...optionsErr } : null
        callback(err, { ...serviceValues, ...optionsValues })
      })
    })
  }

  handleSubmit = async (e) => {
    e.preventDefault()

    this.setState({ submitting: true, formErrorMessage: false })

    this.validatedFormsFields(async (err, values) => {
      if (err) {
        this.setState({ submitting: false, formErrorMessage: 'Validation failed' })
        return
      }
      try {
        const service = await (await KoreApi.client()).UpdateService(this.props.team.metadata.name, values.serviceName, this.generateServiceResource(values))
        message.loading('Cluster application service requested...')

        return this.props.handleSubmit(service)
      } catch (err) {
        console.error('Error saving application service', err)
        this.setState({
          submitting: false,
          formErrorMessage: (err.fieldErrors && err.message) ? err.message : 'An error occurred requesting the application service, please try again',
          validationErrors: err.fieldErrors // This will be undefined on non-validation errors, which is fine.
        })
      }
    })
  }

  handleSelectKind = (kind) => {
    this.setState({
      selectedServiceKind: kind,
      servicePlanOverride: null,
      validationErrors: null
    })
  }

  handleServicePlanOverride = servicePlanOverrides => {
    this.setState({ servicePlanOverride: servicePlanOverrides })
  }

  disableButton = () => {
    if (!this.state.selectedServiceKind) {
      return true
    }
    return false
  }

  cancel = () => {
    this.props.form.resetFields()
    this.setState({ ...ClusterApplicationServiceForm.initialState })
    this.props.handleCancel()
  }

  render() {
    if (this.state.dataLoading || !this.props.team) {
      return <Icon type="loading" />
    }
    const formConfig = {
      layout: 'horizontal',
      labelAlign: 'left',
      hideRequiredMark: true,
      labelCol: {
        sm: { span: 24 },
        md: { span: 24 },
        lg: { span: 6 }
      },
      wrapperCol: {
        sm: { span: 24 },
        md: { span: 24 },
        lg: { span: 18 }
      }
    }

    const { getFieldDecorator } = this.props.form
    const { serviceKinds, selectedServiceKind, namespaceClaims, formErrorMessage, submitting } = this.state

    let filteredServicePlans = []
    if (selectedServiceKind) {
      filteredServicePlans = this.state.servicePlans.items.filter(p => p.spec.kind === selectedServiceKind)
    }

    return (
      <Form {...formConfig} onSubmit={this.handleSubmit}>
        <FormErrorMessage message={formErrorMessage} />
        <Card style={{ marginBottom: '20px' }}>
          <Alert
            message="Cluster application service"
            description="Select the service you would like to use."
            type="info"
            showIcon
            style={{ marginBottom: '20px' }}
          />

          <Form.Item label="Service type">
            {getFieldDecorator('serviceKind', {
              rules: [{ required: true, message: 'Please select your service type!' }],
            })(
              <Select
                onChange={this.handleSelectKind}
                placeholder="Choose service type"
                showSearch
                optionFilterProp="children"
                filterOption={(input, option) =>
                  option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
                }
              >
                {serviceKinds.items.map(k => <Select.Option key={k.metadata.name} value={k.metadata.name}>{k.spec.displayName || k.metadata.name}</Select.Option>)}
              </Select>
            )}
          </Form.Item>

          {selectedServiceKind && (
            <ServiceOptionsForm
              team={this.props.team}
              selectedServiceKind={selectedServiceKind}
              servicePlans={filteredServicePlans}
              teamServices={this.props.teamServices}
              onServicePlanOverridden={this.handleServicePlanOverride}
              validationErrors={this.state.validationErrors}
              wrappedComponentRef={inst => this.serviceOptionsForm = inst}
            />
          )}
        </Card>

        <Card>
          <Alert
            message="Target namespace"
            description="Select the namespace you would like service to be deploying into."
            type="info"
            showIcon
            style={{ marginBottom: '20px' }}
          />
          <Form.Item label="Namespace">
            {getFieldDecorator('namespace', {
              rules: [{ required: true, message: 'Please select the target namespace!' }],
              initialValue: namespaceClaims.items.length === 1 ? namespaceClaims.items[0].spec.name : undefined,
            })(
              <Select placeholder="Choose target namespace">
                {namespaceClaims.items.map(nc => <Select.Option key={nc.spec.name} value={nc.spec.name}>{nc.spec.name}</Select.Option>)}
              </Select>
            )}
          </Form.Item>
        </Card>

        <Form.Item style={{ marginTop: '20px', marginBottom: 0 }}>
          <Button type="primary" htmlType="submit" loading={submitting} disabled={this.disableButton()}>{this.state.submitButtonText}</Button>
          <Button type="link" onClick={this.cancel}>Cancel</Button>
        </Form.Item>
      </Form>
    )
  }
}

const WrappedClusterApplicationServiceForm = Form.create({ name: 'cluster_application' })(ClusterApplicationServiceForm)

export default WrappedClusterApplicationServiceForm
