import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';
import { toNft } from './src/sidebars/technical-docs/to-nft';
import { sgroups } from './src/sidebars/technical-docs/sgroups';
import { terraformItems } from './src/sidebars/users-docs/terraform.items';

const sidebars: SidebarsConfig = {

  informationSidebar: [
    {
      type: 'doc',
      label: 'Введение',
      id: 'info/introduction',
    },
    {
      type: 'doc',
      label: 'Выбор инструмента',
      id: 'info/toolset'
    },
    {
      type: 'category',
      label: 'Архитектура',
      items: [{
        type: 'doc',
        label: 'Абстракции',
        id: 'info/architecture/abstractions'
      },
      {
        type: 'doc',
        label: 'Компоненты',
        id: 'info/architecture/components'
      }]
    },
    {
      type: 'doc',
      label: 'Терминология',
      id: 'terminology'
    }
  ],

  UserDocs: [{
    type: 'doc',
    label: 'Terraform module',
    id: 'user/swarm-spec'
  },
  {
    type: 'category',
    label: 'Terraform',
    items: terraformItems
  }],

  techDocs: [{
    type: 'category',
    label: 'to-nft',
    items: toNft
  },
  {
    type: 'category',
    label: 'sgroups',
    items: sgroups
  }]
};

export default sidebars;
