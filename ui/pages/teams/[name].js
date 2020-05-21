import React from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import Link from 'next/link'
import Router from 'next/router'
import Error from 'next/error'
import { Typography, Card, List, Button, message, Drawer, Badge, Alert, Icon, Modal, Dropdown, Menu, Tabs, Divider } from 'antd'
const { Paragraph, Text } = Typography
const { TabPane } = Tabs
import getConfig from 'next/config'
const { publicRuntimeConfig } = getConfig()

import Breadcrumb from '../../lib/components/layout/Breadcrumb'
import Cluster from '../../lib/components/teams/cluster/Cluster'
import Service from '../../lib/components/teams/service/Service'
import NamespaceClaim from '../../lib/components/teams/namespace/NamespaceClaim'
import NamespaceClaimForm from '../../lib/components/teams/namespace/NamespaceClaimForm'
import apiRequest from '../../lib/utils/api-request'
import copy from '../../lib/utils/object-copy'
import apiPaths from '../../lib/utils/api-paths'
import redirect from '../../lib/utils/redirect'
import KoreApi from '../../lib/kore-api'
import MembersTab from '../../lib/components/teams/members/MembersTab'

class TeamDashboard extends React.Component {
  static propTypes = {
    invitation: PropTypes.bool,
    team: PropTypes.object.isRequired,
    user: PropTypes.object.isRequired,
    clusters: PropTypes.object.isRequired,
    services: PropTypes.object.isRequired,
    namespaceClaims: PropTypes.object.isRequired,
    available: PropTypes.object.isRequired,
    teamRemoved: PropTypes.func.isRequired
  }

  static staticProps = {
    title: 'Team dashboard'
  }

  constructor(props) {
    super(props)
    this.state = {
      tabActiveKey: 'clusters',
      clusters: props.clusters,
      memberCount: -1,
      services: props.services,
      createNamespace: false,
      namespaceClaims: props.namespaceClaims
    }
  }

  static async getTeamDetails(ctx) {
    const name = ctx.query.name
    const api = await KoreApi.client(ctx)
    const getTeam = () => api.GetTeam(name)
    const getTeamClusters = () => api.ListClusters(name)
    const getTeamServices = () => publicRuntimeConfig.featureGates['services'] ? api.ListServices(name) : {}
    const getNamespaceClaims = () => api.ListNamespaces(name)
    const getAvailable = () => api.ListAllocations(name, true)

    return axios.all([getTeam(), getTeamClusters(), getTeamServices(), getNamespaceClaims(), getAvailable()])
      .then(axios.spread(function (team, clusters, services, namespaceClaims, available) {
        return { team, clusters, services, namespaceClaims, available }
      }))
      .catch(err => {
        throw new Error(err.message)
      })
  }

  static getInitialProps = async ctx => {
    const teamDetails = await TeamDashboard.getTeamDetails(ctx)
    if (Object.keys(teamDetails.team).length === 0 && ctx.res) {
      /* eslint-disable-next-line require-atomic-updates */
      ctx.res.statusCode = 404
    }
    if (ctx.query.invitation === 'true') {
      teamDetails.invitation = true
    }
    return teamDetails
  }

  componentDidUpdate(prevProps) {
    const teamFound = Object.keys(this.props.team).length
    const prevTeamName = prevProps.team.metadata && prevProps.team.metadata.name
    if (teamFound && this.props.team.metadata.name !== prevTeamName) {
      const state = copy(this.state)
      state.tabActiveKey = 'clusters'
      state.clusters = this.props.clusters
      state.services = this.props.services
      state.namespaceClaims = this.props.namespaceClaims
      this.setState(state)
    }
  }

  handleResourceUpdated = resourceType => {
    return (updatedResource, done) => {
      const state = copy(this.state)
      const resource = state[resourceType].items.find(r => r.metadata.name === updatedResource.metadata.name)
      resource.status = updatedResource.status
      this.setState(state, done)
    }
  }

  handleResourceDeleted = resourceType => {
    return (name, done) => {
      const state = copy(this.state)
      const resource = state[resourceType].items.find(r => r.metadata.name === name)
      resource.deleted = true
      this.setState(state, done)
    }
  }

  deleteCluster = async (name, done) => {
    const team = this.props.team.metadata.name
    try {
      const state = copy(this.state)
      const cluster = state.clusters.items.find(c => c.metadata.name === name)
      await apiRequest(null, 'delete', `${apiPaths.team(team).clusters}/${cluster.metadata.name}`)
      cluster.status.status = 'Deleting'
      cluster.metadata.deletionTimestamp = new Date()
      this.setState(state, done)
      message.loading(`Cluster deletion requested: ${cluster.metadata.name}`)
    } catch (err) {
      console.error('Error deleting cluster', err)
      message.error('Error deleting cluster, please try again.')
    }
  }

  deleteService = async (name, done) => {
    const team = this.props.team.metadata.name
    try {
      const state = copy(this.state)
      const service = state.services.items.find(s => s.metadata.name === name)
      await apiRequest(null, 'delete', `${apiPaths.team(team).services}/${service.metadata.name}`)
      service.status.status = 'Deleting'
      service.metadata.deletionTimestamp = new Date()
      this.setState(state, done)
      message.loading(`Service deletion requested: ${service.metadata.name}`)
    } catch (err) {
      console.error('Error deleting service', err)
      message.error('Error deleting service, please try again.')
    }
  }

  createNamespace = value => {
    return () => {
      const state = copy(this.state)
      state.createNamespace = value
      this.setState(state)
    }
  }

  handleNamespaceCreated = namespaceClaim => {
    const state = copy(this.state)
    state.createNamespace = false
    state.namespaceClaims.items.push(namespaceClaim)
    this.setState(state)
    message.loading(`Namespace "${namespaceClaim.spec.name}" requested on cluster "${namespaceClaim.spec.cluster.name}"`)
  }

  deleteNamespace = async (name, done) => {
    const team = this.props.team.metadata.name
    try {
      const state = copy(this.state)
      const namespaceClaim = state.namespaceClaims.items.find(nc => nc.metadata.name === name)
      await apiRequest(null, 'delete', `${apiPaths.team(team).namespaceClaims}/${name}`)
      namespaceClaim.status.status = 'Deleting'
      namespaceClaim.metadata.deletionTimestamp = new Date()
      this.setState(state, done)
      message.loading(`Namespace deletion requested: ${namespaceClaim.spec.name}`)
    } catch (err) {
      console.error('Error deleting namespace', err)
      message.error('Error deleting namespace, please try again.')
    }
  }

  clusterAccess = async () => {
    const apiUrl = new URL(publicRuntimeConfig.koreApiPublicUrl)

    const profileConfigureCommand = `kore profile configure ${apiUrl.hostname}`
    const loginCommand = 'kore login'
    const kubeconfigCommand = `kore kubeconfig -t ${this.props.team.metadata.name}`

    const InfoItem = ({ num, title }) => (
      <div style={{ marginBottom: '10px' }}>
        <Badge style={{ backgroundColor: '#1890ff', marginRight: '10px' }} count={num} />
        <Text strong>{title}</Text>
      </div>
    )
    Modal.info({
      title: 'Cluster access',
      content: (
        <div style={{ marginTop: '20px' }}>
          <InfoItem num="1" title="Download" />
          <Paragraph>If you haven&apos;t already, download the CLI from <a href="https://github.com/appvia/kore/releases">https://github.com/appvia/kore/releases</a></Paragraph>

          <InfoItem num="2" title="Setup profile" />
          <Paragraph>Create a profile</Paragraph>
          <Paragraph className="copy-command" style={{ marginRight: '40px' }} copyable>{profileConfigureCommand}</Paragraph>
          <Paragraph>Enter the Kore API URL as follows</Paragraph>
          <Paragraph className="copy-command" style={{ marginRight: '40px' }} copyable>{apiUrl.origin}</Paragraph>

          <InfoItem num="3" title="Login" />
          <Paragraph>Login to the CLI</Paragraph>
          <Paragraph className="copy-command" style={{ marginRight: '40px' }} copyable>{loginCommand}</Paragraph>

          <InfoItem num="4" title="Setup access" />
          <Paragraph>Then, you can use the Kore CLI to setup access to your team&apos;s clusters</Paragraph>
          <Paragraph className="copy-command" style={{ marginRight: '40px' }} copyable>{kubeconfigCommand}</Paragraph>
          <Paragraph>This will add local kubernetes configuration to allow you to use <Text
            style={{ fontFamily: 'monospace' }}>kubectl</Text> to talk to the provisioned cluster(s).</Paragraph>
          <Paragraph>See examples: <a href="https://kubernetes.io/docs/reference/kubectl/overview/" target="_blank" rel="noopener noreferrer">https://kubernetes.io/docs/reference/kubectl/overview/</a></Paragraph>
        </div>
      ),
      width: 700,
      onOk() {}
    })
  }

  deleteTeam = async () => {
    try {
      const team = this.props.team.metadata.name
      const api = await KoreApi.client()
      await api.RemoveTeam(team)
      this.props.teamRemoved(team)
      message.success(`Team "${team}" deleted`)
      return redirect({ router: Router, path: '/' })
    } catch (err) {
      console.log('Error deleting team', err)
      message.error('Team could not be deleted, please try again later')
    }
  }

  deleteTeamConfirm = () => {
    const { clusters } = this.state
    if (clusters.items.length > 0) {
      return Modal.warning({
        title: 'Warning: team cannot be deleted',
        content: (
          <>
            <Paragraph strong>The clusters must be deleted first</Paragraph>
            <List
              size="small"
              dataSource={clusters.items}
              renderItem={c => <List.Item>{c.spec.kind} <Text style={{ fontFamily: 'monospace', marginLeft: '15px' }}>{c.metadata.name}</Text></List.Item>}
            />
          </>
        ),
        onOk() {}
      })
    }

    Modal.confirm({
      title: 'Are you sure you want to delete this team?',
      content: 'This cannot be undone',
      okText: 'Yes',
      okType: 'danger',
      cancelText: 'No',
      onOk: this.deleteTeam
    })
  }

  settingsMenu = ({ team }) => {
    const menu = (
      <Menu>
        <Menu.Item key="audit">
          <Link href="/teams/[name]/audit" as={`/teams/${team.metadata.name}/audit`}>
            <a>
              <Icon type="table" style={{ marginRight: '5px' }} />
              Team audit viewer
            </a>
          </Link>
        </Menu.Item>
        <Menu.Item key="security">
          <Link href="/teams/[name]/security" as={`/teams/${team.metadata.name}/security`}>
            <a>
              <Icon type="lock" style={{ marginRight: '5px' }} />
              Team security overview
            </a>
          </Link>
        </Menu.Item>
        <Menu.Item key="delete" className="ant-btn-danger" onClick={this.deleteTeamConfirm}>
          <Icon type="delete" style={{ marginRight: '5px' }} />
          Delete team
        </Menu.Item>
      </Menu>
    )
    return (
      <Dropdown overlay={menu}>
        <Button>
          <Icon type="setting" style={{ marginRight: '10px' }} />
          <Icon type="down" />
        </Button>
      </Dropdown>
    )
  }

  clustersTabContent = () => {
    const team = this.props.team
    const { clusters, namespaceClaims, createNamespace } = this.state
    const hasActiveClusters = Boolean(clusters.items.filter(c => c.status && c.status.status === 'Success').length)

    return (
      <>
        <Card
          title={<div><Text style={{marginRight: '10px'}}>Clusters</Text><Badge style={{backgroundColor: '#1890ff'}}
                                                                                count={clusters.items.filter(c => !c.deleted).length}/>
          </div>}
          style={{marginBottom: '20px'}}
          extra={
            <div>
              {hasActiveClusters ?
                <Text style={{marginRight: '20px'}}><a onClick={this.clusterAccess}><Icon type="eye" theme="twoTone"/>
                  Access</a></Text> :
                null
              }
              <Button type="primary">
                <Link href="/teams/[name]/clusters/new" as={`/teams/${team.metadata.name}/clusters/new`}>
                  <a>+ New</a>
                </Link>
              </Button>
            </div>
          }
        >
          <List
            dataSource={clusters.items}
            renderItem={cluster => {
              const namespaceClaims = (this.state.namespaceClaims.items || []).filter(nc => nc.spec.cluster.name === cluster.metadata.name && !nc.deleted)
              return (
                <Cluster
                  team={team.metadata.name}
                  cluster={cluster}
                  namespaceClaims={namespaceClaims}
                  deleteCluster={this.deleteCluster}
                  handleUpdate={this.handleResourceUpdated('clusters')}
                  handleDelete={this.handleResourceDeleted('clusters')}
                  refreshMs={10000}
                  propsResourceDataKey="cluster"
                  resourceApiPath={`${apiPaths.team(team.metadata.name).clusters}/${cluster.metadata.name}`}
                />
              )
            }}
          >
          </List>
        </Card>
        <Card
          title={<div><Text style={{ marginRight: '10px' }}>Namespaces</Text><Badge style={{ backgroundColor: '#1890ff' }} count={namespaceClaims.items.filter(c => !c.deleted).length} /></div>}
          style={{ marginBottom: '20px' }}
          extra={clusters.items.length > 0 ? <Button type="primary" onClick={this.createNamespace(true)}>+ New</Button> : null}
        >
          <List
            dataSource={namespaceClaims.items}
            renderItem={namespaceClaim =>
              <NamespaceClaim
                team={team.metadata.name}
                namespaceClaim={namespaceClaim}
                deleteNamespace={this.deleteNamespace}
                handleUpdate={this.handleResourceUpdated('namespaceClaims')}
                handleDelete={this.handleResourceDeleted('namespaceClaims')}
                refreshMs={15000}
                propsResourceDataKey="namespaceClaim"
                resourceApiPath={`${apiPaths.team(team.metadata.name).namespaceClaims}/${namespaceClaim.metadata.name}`}
              />
            }
          >
          </List>
        </Card>
        <Drawer
          title="Create namespace"
          placement="right"
          closable={false}
          onClose={this.createNamespace(false)}
          visible={createNamespace}
          width={700}
        >
          <NamespaceClaimForm team={team.metadata.name} clusters={clusters} handleSubmit={this.handleNamespaceCreated} handleCancel={this.createNamespace(false)}/>
        </Drawer>
      </>
    )
  }
  getTabTitle = ({ title, count, icon }) => (
    <span>
      {title}
      {count !== undefined && count !== -1 && <Badge showZero={true} style={{ marginLeft: '10px', backgroundColor: '#1890ff' }} count={count} />}
      {icon}
    </span>
  )

  render() {
    const { team, invitation } = this.props

    if (Object.keys(team).length === 0) {
      return <Error statusCode={404} />
    }

    const { membersCount, namespaceClaims, clusters, services, tabActiveKey } = this.state

    return (
      <div>
        <div style={{ display: 'inline-block', width: '100%' }}>
          <div style={{ float: 'left', marginTop: '8px' }}>
            <Breadcrumb items={[{ text: team.spec.summary }]} />
          </div>
          <div style={{ float: 'right' }}>
            <this.settingsMenu team={team} />
          </div>
        </div>
        <Paragraph>
          <Text strong>{team.spec.description}</Text>
          <Text style={{ float: 'right' }}><Text strong>Team ID: </Text>{team.metadata.name}</Text>
        </Paragraph>
        {invitation ? (
          <Alert
            message="You have joined this team from an invitation"
            type="info"
            showIcon
            style={{ marginBottom: '20px' }}
          />
        ) : null}

        <Tabs activeKey={tabActiveKey} onChange={(key) => this.setState({ tabActiveKey: key })}>
          <TabPane
            key="clusters"
            tab={
              <span>
                Clusters
                <Badge showZero={true} style={{ marginLeft: '10px', backgroundColor: '#1890ff' }} count={clusters.items.filter(c => !c.deleted).length} />
              </span>
            }>
            <this.clustersTabContent />
          </TabPane>

          <TabPane key="members" tab={this.getTabTitle({ title: 'Members', count: this.state.memberCount })} forceRender={true}>
            <MembersTab user={this.props.user} team={this.props.team} getMemberCount={(count) => this.setState({ memberCount: count })} />
          </TabPane>
          </TabPane>
        </Tabs>

        <Divider />

        {publicRuntimeConfig.featureGates['services'] ? (
          <Card
            title={<div><Text style={{ marginRight: '10px' }}>Services</Text><Badge style={{ backgroundColor: '#1890ff' }} count={services.items.filter(c => !c.deleted).length} /></div>}
            style={{ marginBottom: '20px' }}
            extra={
              <div>
                <Button type="primary">
                  <Link href="/teams/[name]/services/new" as={`/teams/${team.metadata.name}/services/new`}>
                    <a>+ New</a>
                  </Link>
                </Button>
              </div>
            }
          >
            <List
              dataSource={services.items}
              renderItem={service => {
                return (
                  <Service
                    team={team.metadata.name}
                    service={service}
                    namespaceClaims={namespaceClaims}
                    deleteService={this.deleteService}
                    handleUpdate={this.handleResourceUpdated('services')}
                    handleDelete={this.handleResourceDeleted('services')}
                    refreshMs={10000}
                    propsResourceDataKey="service"
                    resourceApiPath={`${apiPaths.team(team.metadata.name).services}/${service.metadata.name}`}
                  />
                )
              }}
            >
            </List>
          </Card>
        ): null}
      </div>
    )
  }
}

export default TeamDashboard
