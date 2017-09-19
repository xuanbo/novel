<template>
  <div>
    <top></top>
    <div>
      <Row style="margin-top: 200px">
        <Col span="8" offset="8">
          <Input v-model="word" :placeholder="tip" size="large" @on-enter="search">
              <Select v-model="source" slot="prepend" style="width: 80px">
                <Option value='1'>笔趣阁</Option>
              </Select>
          </Input>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
import top from '@/components/top'

export default {
  name: 'index',
  data () {
    return {
      tip: '搜索一下，你就知道！',
      word: '',
      source: '1'
    }
  },
  components: {
    top: top
  },
  methods: {
    search: function () {
      this.$router.push({name: 'novelSearch', params: { q: this.word }})
    }
  }
}
</script>
