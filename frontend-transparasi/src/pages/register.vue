<template>
  <v-container>
    <div class="d-flex">

      <div class="header-login form-login">
        <v-card class="pa-5" rounded="lg">
          <div class="d-flex align-center">
            <img src="../assets/logo-dprd-1.png" alt="" class="logo-login">
            <div class="">
            <h1 class="title ">Selamat Datang!</h1>
            <p class="caption mb-5">Lanjutkan untuk masuk!</p>
          </div>

          </div>
  <v-form class="form mt-" @submit.prevent="submitLogin()">
    <p class="label-form mb-1">Nama</p>
      <v-text-field
      variant="outlined"
      v-model="input.username">
    </v-text-field>
    <p class="label-form mb-1">Email</p>
      <v-text-field
      variant="outlined"
      v-model="input.email">
    </v-text-field>
      <p class="label-form mb-1">Kata Sandi</p>
      <v-text-field
      variant="outlined"
      :type="showPassword1 ? 'text' : 'password'"
      :append-inner-icon="showPassword1 ? 'mdi-eye' : 'mdi-eye-closed'"
      @click:append-inner="togglePasswordVisibility(1)"
      :rules="passwordRules"
      v-model="input.password">
    </v-text-field>
    <p class="label-form mb-1">Ulangi Kata Sandi</p>
    <v-text-field
      variant="outlined"
      :type="showPassword ? 'text' : 'password'"
      :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-closed'"
      @click:append-inner="togglePasswordVisibility(2)"
      :rules="confirmPasswordRules"
      v-model="passwordRepeat">
    </v-text-field>
    <v-btn class="d-flex justify-space-between button-login" width="100%" append-icon="mdi-chevron-right" color="grey" type="submit" :disabled="!isFormValid">
      Buat Akun
    </v-btn>
  </v-form>
  <div class="d-flex justify-center mt-5">
    <p class="caption">Sudah punya akun? Lanjutkan Untuk<a href="/login"><b> Masuk</b></a></p>
  </div>
        </v-card>
  </div>
    </div>
  </v-container>
  </template>
  <script>
  import axios from 'axios';
  export default{
    data() {
      return {
        passwordRules: [
        v => !!v || 'Kata sandi harus diisi',
        v => v.length >= 6 || 'Kata sandi minimal 6 karakter'
      ],
      confirmPasswordRules: [
        v => !!v || 'Ulangi kata sandi harus diisi',
        v => v === this.input.password || 'Kata sandi tidak cocok'
      ],
        passwordRepeat:'',
        input :{
          username:'',
          email:'',
          password:'',
        },
        showPassword: false,
        showPassword1:false // State untuk visibilitas password
      }
    },
    mounted() {

    },
    methods:{
      submitLogin(){
        axios.post('/register', this.input)
        .then((res)=>{
          this.$router.push('/login')
            this.$swal({
                toast: "true",
                timer:4000,
                position:"top-end",
                icon: "success",
                title: "Success Register",
                showConfirmButton :false
                });
        })
        .catch(err=>{

            this.$swal({
            toast: "true",
            timer:2000,
            position:"top-end",
            icon: "error",
            title: "Error Register",
            text:"Periksa Kelengkapan Data Kembali",
            showConfirmButton :false
            });

        })
      },
      togglePasswordVisibility(id) {
        if (id==2) {
          this.showPassword = !this.showPassword;
        }
        else if(id==1){
          this.showPassword1 = !this.showPassword1
        }
       // Toggle visibility
      },
    },
    computed: {
    passwordMatchError() {
      return this.passwordRepeat && this.passwordRepeat !== this.input.password
        ? ['Kata sandi tidak cocok']
        : []
    },
    isFormValid() {
      return this.input.username &&
             this.input.email &&
             this.input.password &&
             this.passwordRepeat &&
             this.input.password === this.passwordRepeat
    }
  },
  }
  </script>
  <style scooped>

  </style>
