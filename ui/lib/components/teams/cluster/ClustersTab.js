import React from 'react'
import PropTypes from 'prop-types'
import moment from 'moment'
import Link from 'next/link'
import { Button, Col, Divider, Icon, message, Row, Tag, Tooltip, Typography } from 'antd'
const { Paragraph, Text } = Typography

import Cluster from './Cluster'
import ClusterAccessInfo from './ClusterAccessInfo'
import KoreApi from '../../../kore-api'
import copy from '../../../utils/object-copy'
import { inProgressStatusList, statusColorMap, statusIconMap } from '../../../utils/ui-helpers'

class ClustersTab extends React.Component {

  static propTypes = {
    team: PropTypes.object.isRequired,
    getClusterCount: PropTypes.func
  }

  state = {
    dataLoading: true,
    clusters: [],
    namespaceClaims: [],
    plans: [],
    revealNamespaces: {},
    createNamespace: false
  }

  async fetchComponentData () {
    try {
      const team = this.props.team.metadata.name
      const api = await KoreApi.client()
      let [ clusters, namespaceClaims, plans ] = await Promise.all([
        api.ListClusters(team),
        api.ListNamespaces(team),
        api.ListPlans()
      ])
      clusters = clusters.items
      namespaceClaims = namespaceClaims.items
      plans = plans.items
      const revealNamespaces = {}
      clusters.filter(cluster => namespaceClaims.filter(nc => nc.spec.cluster.name === cluster.metadata.name).length > 0).forEach(cluster => revealNamespaces[cluster.metadata.name] = true)
      this.props.getClusterCount && this.props.getClusterCount(clusters.length)
      return { clusters, namespaceClaims, plans, revealNamespaces }
    } catch (err) {
      console.error('Unable to load data for clusters tab', err)
      return {}
    }
  }

  componentDidMount() {
    return this.fetchComponentData().then(data => {
      this.setState({ ...data, dataLoading: false })
    })
  }

  componentDidUpdate(prevProps) {
    if (prevProps.team.metadata.name !== this.props.team.metadata.name) {
      this.setState({ dataLoading: true, clusters: [] })
      return this.fetchComponentData().then(data => this.setState({ ...data, dataLoading: false }))
    }
  }

  handleResourceUpdated = resourceType => {
    return (updatedResource, done) => {
      const resourceList = copy(this.state[resourceType])
      const resource = resourceList.find(r => r.metadata.name === updatedResource.metadata.name)
      resource.status = updatedResource.status
      this.setState({ [resourceType]: resourceList }, done)
    }
  }

  handleResourceDeleted = resourceType => {
    return (name, done) => {
      const resourceList = copy(this.state[resourceType])
      const resource = resourceList.find(r => r.metadata.name === name)
      resource.deleted = true

      const revealNamespaces = copy(this.state.revealNamespaces)
      if (resourceType === 'namespaceClaims') {
        revealNamespaces[resource.spec.cluster.name] = Boolean(resourceList.filter(nc => !nc.deleted && nc.spec.cluster.name === resource.spec.cluster.name).length)
      }

      this.setState({ [resourceType]: resourceList, revealNamespaces }, () => {
        this.props.getClusterCount && this.props.getClusterCount(this.state.clusters.filter(c => !c.deleted).length)
        done()
      })
    }
  }

  deleteCluster = async (name, done) => {
    const team = this.props.team.metadata.name
    try {
      const clusters = copy(this.state.clusters)
      const cluster = clusters.find(c => c.metadata.name === name)
      await (await KoreApi.client()).RemoveCluster(team, cluster.metadata.name)
      cluster.status.status = 'Deleting'
      cluster.metadata.deletionTimestamp = new Date()
      this.setState({ clusters }, done)
      message.loading(`Cluster deletion requested: ${cluster.metadata.name}`)
    } catch (err) {
      console.error('Error deleting cluster', err)
      message.error('Error deleting cluster, please try again.')
    }
  }

  clusterNamespaceList = ({ namespaceClaims }) => {
    return (
      <Row style={{ marginLeft: '50px' }}>
        <Col>
          <Text strong style={{ marginRight: '8px' }}>Namespaces: </Text>
          {namespaceClaims.map(namespaceClaim => {
            const status = namespaceClaim.status.status || 'Pending'
            const created = moment(namespaceClaim.metadata.creationTimestamp).fromNow()
            return (
              <span key={namespaceClaim.metadata.name} style={{ marginRight: '5px' }}>
                <Tooltip title={`Created ${created}`}>
                  <Tag color={statusColorMap[status] || 'red'}>{namespaceClaim.spec.name} {inProgressStatusList.includes(status) ? <Icon type="loading" /> : <Icon type={statusIconMap[status]} />}</Tag>
                </Tooltip>
              </span>
            )
          })}
        </Col>
      </Row>
    )
  }

  render() {
    const { team } = this.props
    const { dataLoading, clusters, namespaceClaims, plans } = this.state

    const hasActiveClusters = Boolean(clusters.filter(c => !c.deleted).length)

    return (
      <>
        <div>
          <Button type="primary">
            <Link href="/teams/[name]/clusters/new" as={`/teams/${team.metadata.name}/clusters/new`}>
              <a>New cluster</a>
            </Link>
          </Button>
          {!dataLoading && hasActiveClusters && <ClusterAccessInfo buttonStyle={{ float: 'right' }} team={this.props.team} />}
        </div>

        <Divider />

        {dataLoading ? (
          <Icon type="loading" />
        ) : (
          <>
            {!hasActiveClusters && <Paragraph type="secondary">No clusters found for this team</Paragraph>}
            {clusters.map((cluster, idx) => {
              const filteredNamespaceClaims = (namespaceClaims || []).filter(nc => nc.spec.cluster.name === cluster.metadata.name)
              const activeNamespaces = filteredNamespaceClaims.filter(nc => !nc.deleted)
              return (
                <React.Fragment key={cluster.metadata.name}>
                  <Cluster
                    team={team.metadata.name}
                    cluster={cluster}
                    plan={plans.find(plan => plan.metadata.name === cluster.spec.plan)}
                    namespaceClaims={activeNamespaces}
                    deleteCluster={this.deleteCluster}
                    handleUpdate={this.handleResourceUpdated('clusters')}
                    handleDelete={this.handleResourceDeleted('clusters')}
                    refreshMs={10000}
                    propsResourceDataKey="cluster"
                    resourceApiPath={`/teams/${team.metadata.name}/clusters/${cluster.metadata.name}`}
                  />
                  {!cluster.deleted && filteredNamespaceClaims.length > 0 && this.clusterNamespaceList({ namespaceClaims: filteredNamespaceClaims })}
                  {!cluster.deleted && idx < clusters.length - 1 && <Divider />}
                </React.Fragment>
              )
            })}
          </>
        )}
      </>
    )
  }
}

export default ClustersTab
