<template>
  <div>
    <top></top>
    <div style="margin-top: 20px">
      <Row>
        <Col span="20" offset="2" style="background:#eee;padding: 20px">
          <Card :bordered="false" >
            <p slot="title">{{ novel.name }} -- {{ novel.author }} -- {{ novel.lastUpdateTime }}</p>
            <p>{{ novel.description }}</p>
          </Card>
        </Col>
      </Row>
        <Row style="margin-top: 10px">
          <Col span="6" v-for="(chapter, index) in novel.chapters" key="chapter.name" style="padding: 5px">
            <Tag color="blue" @click.native="show(index)">{{chapter.name}}</Tag>
          </Col>
        </Row>

        <Modal v-model="showChapter" width="1200">
          <h1>{{ chapter.name }}</h1>
          <div>
            <pre>{{ chapter.content }}</pre>
          </div>
        </Modal>
    </div>
  </div>
</template>

<script>
import top from '@/components/top'

export default {
  name: 'novelIndex',
  data () {
    return {
      url: this.$route.query.url,
      novel: {},
      showChapter: false,
      chapter: {content: ''}
    }
  },
  components: {
    top: top
  },
  mounted () {
    this.fetchNovel()
  },
  methods: {
    fetchNovel: function () {
      let url = '/novel?url=' + this.url
      this.$http.get(url).then(resp => {
        this.novel = resp
      })
    },
    fetchChapter: function (chapterUrl) {
      let url = '/novel/chapter?url=' + chapterUrl
      this.$http.get(url).then(resp => {
        this.chapter = resp
        this.toggleChapter()
      })
    },
    toggleChapter: function () {
      this.showChapter = !this.showChapter
    },
    show (index) {
      this.fetchChapter(this.novel.chapters[index].url)
    }
  }
}
</script>
