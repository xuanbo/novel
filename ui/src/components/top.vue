<template>
  <Menu mode="horizontal" :theme="theme">
      <Menu-item name="1">
          <Icon type="home"></Icon>
          <router-link to="/">首页</router-link>
      </Menu-item>
      <Menu-item name="2">
        <Icon type="person"></Icon>
        <router-link to="/about">关于</router-link>
      </Menu-item>
      <Menu-item name="3">
        <Icon type="edit"></Icon>
        <router-link to="/recommend">建议</router-link>
      </Menu-item>
  </Menu>
</template>

<script>
export default {
  name: 'header',
  data () {
    return {
      theme: 'light'
    }
  }
}
</script>

<style>
  a {
    color: #000000;
  }
</style>