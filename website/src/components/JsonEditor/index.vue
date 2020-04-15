<template>
  <div class="json-editor">
    <v-jsoneditor v-model="json" :options="options" :plus="false" height="400px" @error="onError" />
  </div>
</template>

<script>
import VJsoneditor from 'v-jsoneditor'

export default {
  name: 'JsonEditor',
  components: {
    VJsoneditor
  },
  data() {
    return {
      json: [
        {
          'items': [
            {
              'market_type': 'forexdata',
              'symbol': 'XAUUSD'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'UKOIL'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'CORN'
            }
          ],
          'name': ''
        },
        {
          'items': [
            {
              'market_type': 'forexdata',
              'symbol': 'XAUUSD'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'XAGUSD'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'AUTD'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'AGTD'
            }
          ],
          'name': '贵金属'
        },
        {
          'items': [
            {
              'market_type': 'forexdata',
              'symbol': 'CORN'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'WHEAT'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'SOYBEAN'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'SUGAR'
            }
          ],
          'name': '农产品'
        },
        {
          'items': [
            {
              'market_type': 'forexdata',
              'symbol': 'UKOIL'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'USOIL'
            },
            {
              'market_type': 'forexdata',
              'symbol': 'NGAS'
            }
          ],
          'name': '能源化工'
        }
      ]
    }
  },
  methods: {
    onError() {
      console.log('error')
    }
  }
}
</script>

<style scoped>
.json-editor {
  height: 100%;
  position: relative;
}
</style>
