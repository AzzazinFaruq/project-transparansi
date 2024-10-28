<template>
  <h1 class="">401 Forbidden Access</h1>
  <p class="mb-6">You Cannot Access This Page</p>
  <v-btn class="" @click="back()" color="red">Kembali ke halaman sebelumnya</v-btn>
</template>
<script>
export default{
  data() {
    return {

    }
  },
  methods: {
  back(){
    this.$router.go(-1)
  }
  },
}
</script>
