<template>
  <div class="header">
    <div class="header-top">
      <div><i class="el-icon-phone"></i>全国服务热线:12345678901</div>
      <div>
        <i class="el-icon-eleme top-icon"></i><span>小程序</span>
        <i class="el-icon-eleme top-icon"></i><span>公众号</span>
        <i class="el-icon-mobile-phone top-icon"></i><span>客服</span>
      </div>
    </div>
    <div class="header-menu">
      <div class="header-manu-bar">
        <el-menu
          :default-active="activeIndex"
          text-color="#555"
          class="el-menu-demo"
          mode="horizontal"
          @select="handleSelect"
        >
          <template v-for="(item, index) in menus" :key="item.name">
            <el-menu-item :index="index" :url="item.url" v-if="!item.children || item.children.length === 0">
              {{ item.name }}
            </el-menu-item>
            <el-sub-menu :index="index" v-else>
              <template #title>{{ item.name }}</template>
              <template v-for="c in item.children" :key="c.name">
                <el-menu-item :url="c.url">{{ c.name }}</el-menu-item>
              </template>
            </el-sub-menu>
          </template>
        </el-menu>
      </div>
      <div class="header-login">
        <span class="header-location"
        ><i class="el-icon-location-information"></i>山东</span
        >
        <el-button type="primary" size="mini">登录</el-button>
        <el-button size="mini">注册</el-button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Header",
  props: {
    msg: String,
    activeIndex: {
      type: String,
      default: '0'
    }
  },
  data() {
    return {
      menus: [
        {
          name: "首页",
          url: "/",
        },
        {
          name: "填志愿",
          url: "/",
          children: [
            {
              name: "自动填志愿",
              url: "",
            },
            {
              name: "一键填写",
              url: "",
            },
          ],
        },
        {
          name: "高考资讯",
          url: "/",
        },
        {
          name: "在线课堂",
          url: "/",
        },
        {
          name: "购买志愿卡",
          url: "/",
        },
        {
          name: "联系我们",
          url: "/",
        },
      ],
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
  },
}
</script>

<style scoped>
.header-top {
  background-color: #eee;
  height: 30px;
  font-size: 14px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}

.top-icon {
  margin-left: 10px;
}

.header-menu {
  display: flex;
  border-bottom: 1px solid #22a2c3;
}

.header-login {
  width: 200px;
  font-size: 15px;
  margin-right: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-location {
  margin-right: 10px;
}

.header-manu-bar {
  flex: 1 1 auto;
  text-align: center;
  box-sizing: border-box;
  display: flex;
  justify-content: center;
  font-weight: 600;
}

.el-menu.el-menu--horizontal {
   border-bottom: none !important;
}
</style>
