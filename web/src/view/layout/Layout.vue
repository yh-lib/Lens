<script setup>
import Aside from './components/Aside.vue';
import Header from './components/Header.vue';
import Main from './components/Main.vue';

</script>

<template>
    <el-container class="el-container">
      <Aside style="background: white;"></Aside>
      <el-container class="el-container__el-container">
        <Header style="height: 100px;"></Header>
        <Main></Main>
      </el-container>
    </el-container>
</template>

<style scoped>
.el-container{
  width: 100%;
  height: 100vh;
  .el-container__el-container{
    width: 100%;
    height: 100vh;
    display: flex;
    flex-direction: column;
  }
}
</style>
