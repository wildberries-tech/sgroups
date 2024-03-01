/* eslint-disable import/no-default-export */
import React, { FC } from 'react'
import Layout from '@theme/Layout'
import { Header } from '@site/src/components/molecules'

const LandingPage: FC = () => {
  return (
    <Layout title="SGroups" description="SGroups">
      <Header />
      <div style={{ minHeight: '90vh' }} />
    </Layout>
  )
}

export default LandingPage
