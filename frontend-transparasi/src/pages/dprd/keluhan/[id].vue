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
          <!-- Informasi Keluhan -->
          <div class="mb-5">
            <label class="label-form mb-2 d-block">Nama Pelapor</label>
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

          <!-- Status Badge -->
          <div class="mb-5">
            <v-chip
              :color="status === 'Sudah Ditanggapi' ? '#387144' : '#F29727'"
              style="color: white"
            >
              {{ status }}
            </v-chip>
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
                placeholder="Masukkan tanggapan Anda"
                variant="outlined"
                density="compact"
                rows="4"
                :readonly="status == 'Sudah Ditanggapi'"
              ></v-textarea>
            </div>

            <div class="d-flex gap-2" v-if="status != 'Sudah Ditanggapi'">
              <v-btn
                color="#387144"
                style="color: white"
                prepend-icon="mdi-content-save"
                @click="update()"
              >
                Simpan Tanggapan
              </v-btn>
              <v-btn
                color="#BF3232"
                style="color: white"
                @click="tanggapi = false"
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
import axios from 'axios'
import Swal from 'sweetalert2'

export default {
  data() {
    return {
      username: '',
      program: '',
      keluhan: '',
      status: '',
      input: {
        tanggapan: '',
        user_tanggapan: ''
      },
      tanggapi: false
    }
  },

  mounted() {
    this.getKeluhan()
    this.getUser()
  },

  methods: {
    getUser() {
      axios.get(`/api/user`)
        .then(res => {
          this.input.user_tanggapan = res.data.data.Id
        })
        .catch(error => {
          console.error('Error:', error)
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal mengambil data user',
          })
        })
    },

    getKeluhan() {
      axios.get(`/api/aduan/${this.$route.params.id}`)
        .then(res => {
          const data = res.data.data
          this.username = data.User.username
          this.program = data.Program.nama_program
          this.keluhan = data.keluhan
          this.status = data.status
          this.input.tanggapan = data.tanggapan
        })
        .catch(error => {
          console.error('Error:', error)
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal mengambil data keluhan',
          })
        })
    },

    update() {
      axios.put(`/api/aduan/tanggapi/${this.$route.params.id}`, this.input)
        .then(() => {
          Swal.fire({
            icon: 'success',
            title: 'Berhasil!',
            text: 'Tanggapan berhasil disimpan',
            showConfirmButton: false,
            timer: 1500
          }).then(() => {
            this.$router.go(-1)
          })
        })
        .catch(error => {
          console.error('Error:', error)
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal menyimpan tanggapan',
          })
        })
    },

    back() {
      this.$router.go(-1)
    }
  }
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

.text-body-1 {
  font-size: 16px;
  line-height: 1.5;
  color: rgba(0, 0, 0, 0.87);
}
</style>
