<template>
  <div>
    <v-card class="pa-5">
      <div class="d-flex justify-space-between align-center mb-5">
        <div>
          <p class="text-h5 font-weight-bold">Edit Institusi</p>
          <p class="text-subtitle-1 text-grey">Edit informasi institusi</p>
        </div>
        <v-btn 
          color="error" 
          @click="$router.go(-1)"
          prepend-icon="mdi-arrow-left"
        >
          Kembali
        </v-btn>
      </div>

      <v-divider class="mb-5"></v-divider>

      <v-row>
        <v-col cols="12" md="8">
          <v-form @submit.prevent="submitForm" ref="form">
            <!-- Informasi Dasar -->
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama Institusi</label>
              <v-text-field
                v-model="formData.nama_institusi"
                placeholder="Masukkan nama institusi"
                variant="outlined"
                density="compact"
                :rules="[rules.required]"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Alamat</label>
              <v-textarea
                v-model="formData.Alamat"
                placeholder="Masukkan alamat institusi"
                variant="outlined"
                density="compact"
                rows="3"
                :rules="[rules.required]"
              ></v-textarea>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Email</label>
              <v-text-field
                v-model="formData.Email"
                placeholder="Masukkan email institusi"
                variant="outlined"
                density="compact"
                type="email"
                :rules="[rules.required, rules.email]"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">No. Telepon</label>
              <v-text-field
                v-model="formData.NoTelp"
                placeholder="Masukkan nomor telepon"
                variant="outlined"
                density="compact"
                :rules="[rules.required]"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Logo</label>
              <div class="preview-container mb-2" v-if="preview || formData.Logo">
                <img :src="preview || `${getImageUrl(formData.Logo)}`" alt="Preview Logo" class="preview-image">
              </div>
              <v-file-input
                v-model="logoFile"
                @change="handleFileUpload"
                accept="image/*"
                placeholder="Pilih logo institusi"
                variant="outlined"
                density="compact"
                prepend-icon="false"
                prepend-inner-icon="mdi-image"
              ></v-file-input>
            </div>

            <!-- Tombol Aksi -->
            <div class="d-flex gap-2 mt-5">
              <v-btn
                color="#387144"
                style="color: white"
                type="submit"
                :loading="loading"
                prepend-icon="mdi-content-save"
              >
                Simpan
              </v-btn>
              <v-btn
                color="#BF3232"
                style="color: white"
                @click="$router.go(-1)"
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
import { getImageUrl } from '@/config/foto';
import axios from 'axios'
import Swal from 'sweetalert2'

export default {
  data() {
    return {
      getImageUrl,
      formData: {
        nama_institusi: '',
        Alamat: '',
        Email: '',
        NoTelp: '',
        Logo: ''
       
      },
      logoFile: null,
      preview: null,
      loading: false,
      rules: {
        required: v => !!v || 'Field ini harus diisi',
        email: v => /.+@.+\..+/.test(v) || 'Email harus valid'
      }
    }
  },

  mounted() {
    this.getInstitusi()
  },

  methods: {
    getInstitusi() {
      axios.get(`/api/institusi/${this.$route.params.id}`)
        .then(res => {

          this.formData = res.data.data
          console.log(this.formData)
        })
        .catch(error => {
          console.error('Error:', error)
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal mengambil data institusi',
          })
        })
    },

    handleFileUpload(event) {
      const file = event.target.files[0]
      if (file) {
        this.logoFile = file
        const reader = new FileReader()
        reader.onload = e => {
          this.preview = e.target.result
        }
        reader.readAsDataURL(file)
      }
    },


    async submitForm() {
      try {
        if (!this.$refs.form.validate()) return

        this.loading = true

        const formData = new FormData()
        Object.keys(this.formData).forEach(key => {
          if (key !== 'logo') {
            formData.append(key, this.formData[key])
          }
        })

        if (this.logoFile) {
          formData.append('logo', this.logoFile)
        }

        await axios.put(`/api/institusi/${this.$route.params.id}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        await Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: 'Data institusi berhasil diperbarui',
          showConfirmButton: false,
          timer: 1500
        })

        this.$router.go(-1)

      } catch (error) {
        console.error('Error:', error)
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Gagal memperbarui data institusi',
        })
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
.label-form {
  font-size: 14px;
  color: rgba(0, 0, 0, 0.6);
}

.preview-container {
  width: 200px;
  height: 200px;
  border: 1px solid #ddd;
  border-radius: 4px;
  overflow: hidden;
}

.preview-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.v-btn {
  text-transform: none !important;
}
</style>
