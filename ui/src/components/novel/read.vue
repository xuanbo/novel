<template>
  <div>
    <top id="top"></top>
    <div>
      <Row type="flex" justify="center" style="font-size: 25px">
        <Col >{{ chapter.name }}</Col>
      </Row>
      <Row justify="center">
        <Col>
          <Affix>
            <Button @click="go(0)">第一章</Button>
            <Button @click="go(-1)"><Icon type="chevron-left"></Icon></Button>
            <Button @click="go(1)"><Icon type="chevron-right"></Icon></Button>
            <Button @click="go(-2)">最后一章</Button>
          </Affix>
        </Col>
      </Row>
      <Row justify="center" style="font-size: 20px">
        <Col span="16" offset="4">
          <vue-markdown :source="chapter.content"></vue-markdown>
        </Col>
      </Row>
    </div>
    <BackTop></BackTop>
  </div>
</template>

<script>
import top from '@/components/top'
import VueMarkdown from 'vue-markdown'

export default {
  name: 'novelRead',
  data () {
    return {
      url: this.$route.query.url,
      index: this.$route.query.index,
      active: '',
      novel: {},
      chapter: {}
    }
  },
  components: {
    top,
    VueMarkdown
  },
  mounted () {
    this.fetchNovel()
  },
  filters: {
    newName: function (val) {
      return '1-' + val
    }
  },
  methods: {
    fetchNovel: function () {
      let url = '/novel?url=' + this.url
      this.$http.get(url).then(resp => {
        this.novel = resp
        this.fetchChapter()
      })
    },
    fetchChapter: function () {
      let chapterUrl = this.novel.chapters[this.index].url
      let url = '/novel/chapter?url=' + chapterUrl
      this.$http.get(url).then(resp => {
        this.chapter = resp
      })
    },
    select: function (index) {
      console.log(index)
      this.index = index
      this.fetchChapter()
    },
    go: function (i) {
      this.index = Number(this.index)
      if (i === 0) {
        this.index = 0
      } else if (i === -1) {
        if (this.index - 1 < 0) {
          this.index = 0
        } else {
          this.index = this.index - 1
        }
      } else if (i === 1) {
        if (this.index + 1 > this.novel.chapters.length - 1) {
          this.index = this.novel.chapters.length - 1
        } else {
          this.index = this.index + 1
        }
      } else {
        this.index = this.novel.chapters.length - 1
      }
      this.fetchChapter()
      this.goTop()
    },
    goTop () {
      document.getElementById('top').scrollIntoView()
    }
  }
}
</script>
