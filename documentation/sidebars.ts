import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';
import { toNft } from './src/sidebars/technical-docs/to-nft';
import { sgroups } from './src/sidebars/technical-docs/sgroups';
import { terraformResourceItems } from './src/sidebars/users-docs/terraform.resource.items';
import { terraformModuleItems } from './src/sidebars/users-docs/terraform.module.items';

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
      id: 'info/terminology'
    }
  ],

  UserDocs: [{
    type: 'category',
    label: 'Terraform module',
    link: {
      type: 'doc',
      id: 'user/terraform-module/main'
    },
    items: terraformModuleItems
    
  },
  {
    type: 'category',
    label: 'Terraform resource',
    items: terraformResourceItems
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
