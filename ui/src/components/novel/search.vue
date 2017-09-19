<template>
  <div>
    <top></top>
    <div style="margin-top: 50px">
      <Table border stripe :columns="columns" :data="result"></Table>
    </div>
  </div>
</template>

<script>
import top from '@/components/top'

export default {
  name: 'novelSearch',
  data () {
    return {
      q: this.$route.params.q,
      columns: [
        {
          title: '操作',
          key: 'action',
          width: 100,
          align: 'center',
          render: (h, params) => {
            return h('div', [
              h('Button', {
                props: {
                  type: 'primary',
                  size: 'small'
                },
                style: {
                  marginRight: '5px'
                },
                on: {
                  click: () => {
                    this.show(params.index)
                  }
                }
              }, '浏览')
            ])
          }
        },
        {
          title: '小说名',
          key: 'novelName'
        },
        {
          title: '作者',
          key: 'author'
        },
        {
          title: '最新章节',
          key: 'lastChapterName'
        },
        {
          title: '更新时间',
          key: 'updateTime'
        },
        {
          title: '字数',
          key: 'wordCount'
        },
        {
          title: '状态',
          key: 'status'
        }
      ],
      result: []
    }
  },
  components: {
    top: top
  },
  mounted () {
    this.search()
  },
  methods: {
    search: function () {
      let url = '/novel/search/' + this.q
      this.$http.get(url).then(resp => {
        console.log(resp)
        this.result = resp
      })
    },
    show (index) {
      let url = this.result[index].novelUrl
      this.$router.push({name: 'novelIndex', query: { url: url }})
    }
  }
}
</script>
