<template>
<v-container fluid>
  <div class="d-flex">

    <div class="header-login form-login">
      <v-card class="pa-5" rounded="lg">
        <div class="d-flex align-center">
          <img src="../assets/logo-dprd-1.png" alt="" class="logo-login">
          <div class="">
          <h1 class="title ">Halo</h1>
          <p class="caption mb-5">Lanjutkan untuk masuk!</p>
        </div>

        </div>
    <v-form class="form mt-" @submit.prevent="submitLogin()">
    <p class="label-form mb-1">Email</p>
    <v-text-field
    variant="outlined"
    v-model="input.email">
  </v-text-field>
  <p class="label-form mb-1">Kata Sandi</p>
  <v-text-field
    variant="outlined"
    :type="showPassword ? 'text' : 'password'"
    :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-closed'"
    @click:append-inner="togglePasswordVisibility()"
    v-model="input.password">
  </v-text-field>
  <v-checkbox v-model="input.remember_me" label="Keep me logged in"></v-checkbox>
  <v-btn class="d-flex justify-space-between button-login" width="100%" append-icon="mdi-arrow-right" color="grey" type="submit">
    Masuk
  </v-btn>
  </v-form>
      </v-card>
</div>
  </div>
</v-container>
</template>
<script>
import axios from 'axios';
import { authStore } from '@/store/auth';
import router from '@/router';

const auth = authStore();


export default{
  data() {
    return {
      input :{
        email:'',
        password:'',
        remember_me : false,
      },
      showPassword: false, // State untuk visibilitas password
    }
  },
  mounted() {
    this.check();
  },
  methods:{
    check(){
      axios.get('/api/user')
        .then((res)=>{
          console.log(res.data.status)
          var stat = res.data.status;
          var name = res.data.username;
          if (stat ==  true) {
            this.$swal({
              toast: "true",
              timer:4000,
              position:"top-end",
              icon: "warning",
              title: "Peringatan",
              text:'Anda Sudah Login Sebagai "'+name+'"',
              showConfirmButton :false
              });
            this.$router.push("/dashboard")
          }})
    },
    submitLogin(){
      axios.post('/login', this.input)
      .then((res)=>{
        var name=res.data.username;
        if (res.data.authenticated == true) {
          this.$router.push("/dashboard")
          this.$swal({
              toast: "true",
              timer:4000,
              position:"top-end",
              icon: "success",
              title: "Success Login",
              text:"Welcome "+name,
              showConfirmButton :false
              });
        }
      })
      .catch(err=>{
        if (err.response.data.authenticated == false) {
          this.$swal({
          toast: "true",
          timer:2000,
          position:"top-end",
          icon: "error",
          title: "Error Login",
          text:"Check your email / password again",
          showConfirmButton :false
          });
        }
      })
    },
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword; // Toggle visibility
    },
  }
}
</script>
<style scooped>

</style>
