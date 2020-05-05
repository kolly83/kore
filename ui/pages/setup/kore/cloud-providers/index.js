import React from 'react'
import PropTypes from 'prop-types'
import Router from 'next/router'
import { Typography, Card } from 'antd'
const { Title, Paragraph } = Typography
import getConfig from 'next/config'
const { publicRuntimeConfig } = getConfig()

import redirect from '../../../../lib/utils/redirect'
import copy from '../../../../lib/utils/object-copy'
import apiRequest from '../../../../lib/utils/api-request'
import apiPaths from '../../../../lib/utils/api-paths'
import GKECredentialsForm from '../../../../lib/components/credentials/GKECredentialsForm'
import CloudSelector from '../../../../lib/components/common/CloudSelector'

class ConfigureCloudProvidersPage extends React.Component {
  static propTypes = {
    allTeams: PropTypes.object.isRequired
  }

  static staticProps = {
    title: 'Configure cluster providers',
    hideSider: true,
    adminOnly: true
  }

  state = {
    selectedCloud: ''
  }

  static getInitialProps = async ctx => {
    const allTeams = await apiRequest(ctx, 'get', apiPaths.teams)
    allTeams.items = allTeams.items.filter(t => !publicRuntimeConfig.ignoreTeams.includes(t.metadata.name))
    return { allTeams }
  }

  handleSelectCloud = cloud => {
    if (this.state.selectedCloud !== cloud) {
      const state = copy(this.state)
      state.selectedCloud = cloud
      this.setState(state)
    }
  }

  handleFormSubmit = () => {
    redirect({
      router: Router,
      path: '/setup/kore/complete'
    })
  }

  render() {
    const { selectedCloud } = this.state
    const { allTeams } = this.props

    return (
      <div>
        <Title>Configure Cloud Cluster Provider</Title>
        <Paragraph>Choose your first cloud provider for your clusters, more can be configured later.</Paragraph>
        <div style={{ marginTop: '20px', marginBottom: '20px' }}>
          <CloudSelector selectedCloud={selectedCloud} handleSelectCloud={this.handleSelectCloud} />
        </div>
        { selectedCloud === 'GKE' ? (
          <Card title="Enter GKE credentials" style={{ paddingBottom: '0' }}>
            <GKECredentialsForm team={publicRuntimeConfig.koreAdminTeamName} allTeams={allTeams} handleSubmit={this.handleFormSubmit} saveButtonText="Save & Verify" inlineVerification={true} />
          </Card>
        ) : null }
      </div>
    )
  }
}

export default ConfigureCloudProvidersPage
