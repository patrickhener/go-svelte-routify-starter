
/**
 * @roxi/routify 2.12.4
 * File generated Thu Feb 18 2021 11:21:08 GMT+0100 (Central European Standard Time)
 */

export const __version = "2.12.4"
export const __timestamp = "2021-02-18T10:21:08.796Z"

//buildRoutes
import { buildClientTree } from "@roxi/routify/runtime/buildRoutes"

//imports


//options
export const options = {}

//tree
export const _tree = {
  "root": true,
  "children": [
    {
      "isIndex": true,
      "isPage": true,
      "path": "/index",
      "id": "_index",
      "component": () => import('../src/pages/index.svelte').then(m => m.default)
    },
    {
      "isDir": true,
      "ext": "",
      "children": [
        {
          "isIndex": true,
          "isPage": true,
          "path": "/other/index",
          "id": "_other_index",
          "component": () => import('../src/pages/other/index.svelte').then(m => m.default)
        }
      ],
      "path": "/other"
    }
  ],
  "path": "/"
}


export const {tree, routes} = buildClientTree(_tree)

