<template>
  <div>
    <div class="big">
      <div class="inner">
        <img src="../../assets/404.webp">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login"

}
</script>

<style lang="scss" scoped>
.big {
  width: 100%;
  height: calc(100vh - 220px);
  position: relative;
}

.inner {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 80%;
  max-width: 500px;
}

.inner img{
  width: 100%;
}
</style>