<template>
<v-container fluid>
  <div class="d-flex justify-center align-center">

    <v-form class="form-login" @submit.prevent="submitLogin()">
    <h3 class="text-center mb-3">LOGIN!</h3>
    <v-text-field
    label="Username"
    variant="outlined"
    v-model="input.email">
  </v-text-field>
  <v-text-field
    label="Password"
    variant="outlined"
    v-model="input.password">
  </v-text-field>
  <v-btn color="success" type="submit">Submit</v-btn>
  </v-form>
  </div>
</v-container>
</template>
<script>
import axios from 'axios';

export default{
  data() {
    return {
      input :{
        email:'',
        password:''
      }
    }
  },
  methods:{
    submitLogin(){
      axios.post('/login', this.input)
      .then((res)=>{
        console.log(res.data)
      })
    }
  }
}
</script>
<style scooped>
.form-login{
width: 500px;
position: absolute;
top: 50%;
left: 50%;
transform: translate(-50%,-50%);
@media screen and (max-width:600px) {
  width: 300px;
}
}
</style>
