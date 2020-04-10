<template>
  <div class="errPage-container">
    <el-button class="pan-back-btn" icon="el-icon-arrow-left" @click="back">
      返回
    </el-button>
    <el-row>
      <el-col :span="12">
        <h2 class="text-jumbo text-ginormous">
          Oops! 找不到页面！
        </h2>
        <h5 class="list-unstyled">
          或者你可以
          <router-link to="/dashboard">
            回首页
          </router-link>
        </h5>
      </el-col>
      <el-col :span="12">
        <img
          :src="errGif"
          alt="小女孩的冰激凌融掉了."
          height="428"
          width="313"
        >
      </el-col>
    </el-row>
  </div>
</template>

<script>
import errGif from '@/assets/404_images/404.gif'

export default {
  name: 'Page404',
  data() {
    return {
      errGif: errGif + '?' + +new Date()
    }
  },
  methods: {
    back() {
      if (this.$route.query.noGoBack) {
        this.$router.push({ path: '/dashboard' })
      } else {
        this.$router.go(-1)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.errPage-container {
  width: 800px;
  max-width: 100%;
  margin: 100px auto;

  .pan-back-btn {
    background: #008489;
    color: #fff;
    border: none !important;
  }

  .text-jumbo {
    font-size: 36px;
    font-weight: 700;
    color: #484848;
  }

  .list-unstyled {
    font-size: 20px;

    a {
      color: #008489;
      text-decoration: none;

      &:hover {
        text-decoration: underline;
      }
    }
  }
}
</style>
