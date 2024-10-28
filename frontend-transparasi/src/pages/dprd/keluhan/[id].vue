<template>
  <v-container>
    <div class="">
      <div class="mb-3">
        <label class="label-form ">Nama</label>
      <p class="mt-1">{{ username}}</p>
      </div>
      <div class="mb-3">
        <label class="label-form">Program</label>
      <p class="mt-1">{{ program }}</p>
      </div>
      <div class="mb-3">
      <label class="label-form ">Keluhan</label>
      <p class="mt-1">{{ keluhan }}</p>
      </div>
      <div class="d-flex justify-start" v-if="tanggapi == false && status =='Belum Ditanggapi'">
        <v-btn variant="tonal" style="background-color:#387144;color: white;text-transform: none;" @click="tanggapi = true">
          Tanggapi
        </v-btn>
        <v-btn variant="tonal" class="ml-2" style="background-color:#BF3232;color: white;text-transform: none;" @click="back()">
          Tutup
        </v-btn>
      </div>
      <div class="" v-if="status == 'Belum Ditanggapi' && tanggapi == true">
        <div class="">
        <v-form>
          <label class="label-form ">Tanggapan</label>
          <v-textarea
          class="mt-3"
          variant="outlined"
          v-model="input.tanggapan"
          >
        </v-textarea>
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
      <div class="" v-else-if="status == 'Sudah Ditanggapi'">
        <div class="">
        <v-form>
          <label class="label-form ">Tanggapan</label>
          <v-textarea
           v-model="input.tanggapan"
          class="mt-3"
          variant="outlined"
          >
        </v-textarea>
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
    </div>
  </v-container>
</template>
<script>
import axios from 'axios';
export default{
  data() {
    return {
      username:'',
      program:'',
      keluhan:'',
      status:'',
      input:{
      tanggapan:''
      },
      tanggapi:false
    }
  },
  mounted() {
    this.getKeluhan()
  },
  methods: {
    getKeluhan(){
      axios.get(`/api/aduan/${this.$route.params.id}`)
      .then(res=>{
        console.log(res.data.data)
       this.username = res.data.data.User.username;
       this.program = res.data.data.Program.nama_program
       this.keluhan = res.data.data.keluhan
       this.status = res.data.data.status
       this.input.tanggapan = res.data.data.tanggapan
      })
    },
    update(){
      axios.put(`/api/aduan/tanggapi/${this.$route.params.id}`, this.input)
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
