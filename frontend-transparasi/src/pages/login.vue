<template>
<v-container fluid>
  <div class="d-flex ">
    <div class="header-login form-login">
      <h1 class="title ">Halo</h1>
      <p class="caption mb-5">Lanjutkan untuk masuk!</p>

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
    :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
    @click:append-inner="togglePasswordVisibility()"
    v-model="input.password">
  </v-text-field>
  <v-checkbox  label="Keep me logged in"></v-checkbox>
  <v-btn class="d-flex justify-space-between button-login" width="100%" append-icon="mdi-arrow-right" color="grey" type="submit">
    Masuk
  </v-btn>
  </v-form>
</div>
  </div>
</v-container>
</template>
<script>
import axios from 'axios';
import { authStore } from '@/store/auth';
import router from '@/router';

const auth = authStore();
auth.check(router);

export default{
  data() {
    return {
      input :{
        email:'',
        password:''
      },
      showPassword: false, // State untuk visibilitas password
    }
  },
  methods:{
    submitLogin(){
      axios.post('/login', this.input)
      .then((res)=>{
        console.log(res.data.authenticated)
        if (res.data.authenticated == true) {
          this.$router.push("/dashboard")
        } else {
          alert(res.data.error)
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
.header-login{

  .title{

  }
  .caption{

  }
}
.form-login{
width: 500px;
position: absolute;
top: 50%;
left: 50%;
transform: translate(-50%,-50%);
@media screen and (max-width:600px) {
  width: 300px;
}
.form{
  .label-form{
  font-weight: bold;

}
}


.button-login{
  text-transform: none; /* Menonaktifkan auto capitalization */
  font-size: 16px;
  height: 50px;
}
}
</style>
