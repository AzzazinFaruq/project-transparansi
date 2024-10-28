<template>
  <v-container>
    <div class="">
      <div class="">
        <v-form>
          <label class="label-form">Nama</label>
          <v-text-field v-model="user.username" variant="outlined"></v-text-field>
          <label class="label-form">Email</label>
          <v-text-field v-model="user.email" variant="outlined"></v-text-field>
          <h2 class="my-5">Detail Profile</h2>
          <label class="label-form">Alamat</label>
          <v-text-field v-model="user.alamat" variant="outlined"></v-text-field>
          <label class="label-form">No. HP</label>
          <v-text-field v-model="user.no_hp" variant="outlined"></v-text-field>

          <div class="d-flex justify-start">
            <v-btn variant="tonal" style="background-color:#387144;color: white;text-transform: none;" @click="update()">
              Simpan
            </v-btn>
            <v-btn variant="tonal" class="ml-2" style="background-color:#BF3232;color: white;text-transform: none;" @click="back()">
              Tutup
            </v-btn>
          </div>
        </v-form>
      </div>
    </div>
  </v-container>
</template>
<script>
import axios from 'axios';
export default{
  data() {
    return {
      jabatanItem:[],
      user:{
        username:'',
        email:'',
        password:'',
        no_hp:''

      }
    }
  },
  mounted() {
    this.getUser();
  },
  methods: {
    getUser(){
      axios.get(`/api/user/${this.$route.params.id}`)
      .then(res=>{
        console.log(res.data.data)
        this.user = res.data.data
      })
    },
    update(){
    axios.put(`/api/user/edituser/${this.$route.params.id}`, this.user)
    .then(res=>{
      this.$router.go(-1)
    })
  },
  back(){
    this.$router.go(-1)
  }
  },

}
</script>
