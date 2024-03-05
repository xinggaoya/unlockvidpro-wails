<script setup lang="ts">
import {computed, onMounted, ref} from 'vue'
import {ElMessage, ElNotification} from 'element-plus'
import {DailyQuote} from '../../wailsjs/go/internal/App'
import {WindowFullscreen, WindowUnfullscreen} from "../../wailsjs/runtime";

const playLine = [
  {'name': '2S0', 'url': 'https://jx.2s0.cn/player/?url='},
  {'name': '2ys', 'url': 'https://gj.fenxiangb.com/player/analysis.php?v='},
  {"name": "B站1", "url": "https://jx.jsonplayer.com/player/?url="},
  {'name': '虾米', 'url': 'https://jx.xmflv.com/?url='},
]

const iframeUrl = computed(() => {
  return vipAddress.value + videoAddress.value
})

const vipAddress = ref(playLine[0].url)
const videoAddress = ref('')
const radioValue = ref(playLine[0].name)
const iframeLoading = ref(false)

function inputHande(v: string) {
  localStorage.setItem('videoAddress', v)
}

function handelChange(v: any) {
  vipAddress.value = v.url
  localStorage.setItem('vipAddress', JSON.stringify(v))
  ElMessage.success('切换成功,请稍等片刻...')
  loading()
}

// 加载
function loading() {
  iframeLoading.value = true
  setTimeout(() => {
    iframeLoading.value = false
  }, 1000)
}

// 全屏
function fullScreenHandel() {
  document.addEventListener("fullscreenchange", () => {
    if (document.fullscreenElement) {
      WindowFullscreen();
    } else {
      WindowUnfullscreen();
    }
  });
}

onMounted(() => {
  loading()
  const address = localStorage.getItem('videoAddress')
  const vip = localStorage.getItem('vipAddress')
  if (address) {
    videoAddress.value = address
  }
  if (vip) {
    const data = JSON.parse(vip)
    vipAddress.value = data.url
    radioValue.value = data.name
  }
  fullScreenHandel()
  ElNotification({
    title: '提示',
    message: '部分接口存在恶意广告；仅供测试！！',
    type: 'warning'
  })
  DailyQuote().then((res) => {
    ElNotification({
      title: '每日一句',
      message: res.content,
      type: 'info'
    })
  })
})
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-main>
        <div v-loading="iframeLoading">
          <iframe
              id="playLine"
              :src="iframeUrl"
              width="100%"
              title="YouTube video player"
              allowfullscreen
              frameborder="0"
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          ></iframe>
        </div>
        <div>
          <div>
            <el-form label-position="top">
              <el-form-item label="播放线路">
                <el-radio-group v-model="radioValue">
                  <el-radio
                      :label="item.name"
                      v-for="(item,index) in playLine"
                      @change="handelChange(item)"
                      :key="index"
                      border
                  />
                </el-radio-group>
              </el-form-item>
              <el-form-item label="视频链接">
                <el-input type="textarea" v-model="videoAddress" rows="4" @input="inputHande"
                          placeholder="请输入视频原链接"/>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </el-main>
      <el-footer class="main-center">
        <el-space>
          <el-link type="info" href="https://www.iqiyi.com">爱奇艺</el-link>
          <el-link type="info" href="https://v.qq.com">腾讯视频</el-link>
          <el-link type="info" href="https://www.youku.com">优酷</el-link>
        </el-space>
        <div>
          <el-text tag="mark">请注意部分接口存在恶意广告，请斟酌使用；仅供测试</el-text>
        </div>
      </el-footer>
    </el-container>
  </div>
</template>

<style scoped>
@media screen and (max-width: 768px) {
  .common-layout {
    width: 100%;
    height: 100%;
  }

  #playLine {
    width: 100%;
    height: 250px;
  }
}

@media screen and (min-width: 768px) {
  .common-layout {
    width: 768px;
    height: 100%;
    margin: 0 auto;
  }

  #playLine {
    width: 100%;
    height: 450px;
  }
}

.main-center {
  text-align: center;
}

.main-header {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
