<template>
  <div>
    <v-card class="pa-5" v-if="status == 'Sudah Ditanggapi' || status=='Belum Ditanggapi'">
      <div class="d-flex justify-space-between align-center mb-5">
        <div>
          <p class="text-h5 font-weight-bold">Detail Keluhan</p>
          <p class="text-subtitle-1 text-grey">Informasi lengkap keluhan</p>
        </div>
        <v-btn 
          color="error" 
          @click="back()"
          prepend-icon="mdi-arrow-left"
        >
          Kembali
        </v-btn>
      </div>

      <v-divider class="mb-5"></v-divider>

      <v-row>
        <v-col cols="12" md="8">
          <div class="mb-5">
            <label class="label-form mb-2 d-block">Nama</label>
            <p class="text-body-1">{{ username }}</p>
          </div>

          <div class="mb-5">
            <label class="label-form mb-2 d-block">Program</label>
            <p class="text-body-1">{{ program }}</p>
          </div>

          <div class="mb-5">
            <label class="label-form mb-2 d-block">Keluhan</label>
            <p class="text-body-1">{{ keluhan }}</p>
          </div>

          <!-- Tombol Tanggapi -->
          <div class="d-flex gap-2 mb-5" v-if="tanggapi == false && status =='Belum Ditanggapi'">
            <v-btn
              color="#387144"
              style="color: white"
              prepend-icon="mdi-reply"
              @click="tanggapi = true"
            >
              Tanggapi
            </v-btn>
            <v-btn
              color="#BF3232"
              style="color: white"
              @click="back()"
            >
              Tutup
            </v-btn>
          </div>

          <!-- Form Tanggapan -->
          <v-form v-if="(status == 'Belum Ditanggapi' && tanggapi == true) || status == 'Sudah Ditanggapi'">
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Tanggapan</label>
              <v-textarea
                v-model="input.tanggapan"
                placeholder="Masukkan tanggapan"
                variant="outlined"
                density="comfortable"
                rows="4"
              ></v-textarea>
            </div>

            <div class="d-flex gap-2">
              <v-btn
                color="#387144"
                style="color: white"
                prepend-icon="mdi-content-save"
                @click="update()"
              >
                Simpan
              </v-btn>
              <v-btn
                color="#BF3232"
                style="color: white"
                @click="back()"
              >
                Batal
              </v-btn>
            </div>
          </v-form>
        </v-col>

        <v-col cols="12" md="4">
          <div class="d-flex justify-center">
            <v-img
              src="./logo-dprd-1.png"
              max-width="250"
              contain
              class="mt-5"
            ></v-img>
          </div>
        </v-col>
      </v-row>
    </v-card>
  </div>
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
      tanggapan:'',
      user_tanggapan:''
      },
      tanggapi:false
    }
  },
  mounted() {
    this.getKeluhan()
    this.getUser()
  },
  methods: {
    getUser(){
      axios.get(`/api/user`)
      .then(res=>{
        this.input.user_tanggapan = res.data.data.Id
      })
    },  
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

<style scoped>
.label-form {
  font-size: 14px;
  color: rgba(0, 0, 0, 0.6);
}

.v-btn {
  text-transform: none !important;
}
</style>
