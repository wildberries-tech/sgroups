import { toNft } from './src/sidebars/technical-docs/to-nft'
import { sgroups } from './src/sidebars/technical-docs/sgroups'
import { ruleConfiguration } from './src/sidebars/technical-docs/rule-configuration'

const sidebars = {
  informationSidebar: [
    {
      type: 'doc',
      label: 'Введение',
      id: 'info/introduction',
    },
    {
      type: 'doc',
      label: 'Выбор инструмента',
      id: 'info/toolset',
    },
    {
      type: 'doc',
      label: 'Терминология',
      id: 'info/terminology',
    },
  ],

  techDocs: [
    {
      type: 'doc',
      label: 'Компоненты',
      id: 'tech-docs/components',
    },
    {
      type: 'category',
      label: 'HBF-агент',
      collapsed: false,
      items: toNft,
    },
    {
      type: 'category',
      label: 'HBF-сервер',
      collapsed: false,
      items: sgroups,
    },
    {
      type: 'category',
      label: 'Конфигурация  ресурсов',
      collapsed: false,
      link: {
        type: 'doc',
        id: 'tech-docs/rule-configuration/main',
      },
      items: ruleConfiguration,
    },
    {
      type: 'doc',
      label: 'Требования',
      id: 'tech-docs/installation-system-requirements',
    },
    {
      type: 'doc',
      label: 'Terraform провайдер',
      id: 'tech-docs/installation-terraform',
    },
  ],
}

export default sidebars
