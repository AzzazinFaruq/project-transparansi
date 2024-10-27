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
          <label class="label-form">Jabatan</label>
          <v-text-field v-model="user.jabatan_id" variant="outlined"></v-text-field>
          <label class="label-form">No. HP</label>
          <v-text-field v-model="user.no_hp" variant="outlined"></v-text-field>

          <div class="d-flex justify-start">
            <v-btn variant="tonal" style="background-color:#387144;color: white;text-transform: none;">
              Simpan
            </v-btn>
            <v-btn variant="tonal" class="ml-2" style="background-color:#BF3232;color: white;text-transform: none;">
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
      user:{
        username:'',
        email:'',
        password:'',
        jabatan_id:0,
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
    }
  },
}
</script>
