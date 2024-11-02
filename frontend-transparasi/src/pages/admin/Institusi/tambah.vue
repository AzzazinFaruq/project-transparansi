`<template>
  <div>
    <v-card class="pa-5">
      <div class="d-flex justify-space-between align-center mb-5">
        <div>
          <p class="text-h5 font-weight-bold">Tambah Institusi</p>
          <p class="text-subtitle-1 text-grey">Tambahkan data institusi baru</p>
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

      <v-form @submit.prevent="submitForm" ref="form">
        <v-row>
          <v-col cols="12" md="6">
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
                v-model="formData.alamat"
                placeholder="Masukkan alamat institusi"
                variant="outlined"
                density="compact"
                :rules="[rules.required]"
              ></v-textarea>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Email</label>
              <v-text-field
                v-model="formData.email"
                placeholder="Masukkan email institusi"
                variant="outlined"
                density="compact"
                type="email"
                :rules="[rules.required, rules.email]"
              ></v-text-field>
            </div>
          </v-col>

          <v-col cols="12" md="6">
            <div class="mb-5">
              <label class="label-form mb-2 d-block">No. Telepon</label>
              <v-text-field
                v-model="formData.no_telp"
                placeholder="Masukkan nomor telepon"
                variant="outlined"
                density="compact"
                :rules="[rules.required]"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Logo Institusi</label>
              <v-file-input
                v-model="formData.logo"
                placeholder="Pilih logo institusi"
                variant="outlined"
                density="compact"
                accept="image/*"
                prepend-icon="mdi-camera"
                :rules="[rules.required]"
                @change="previewImage"
              ></v-file-input>
              
              <v-img
                v-if="imagePreview"
                :src="imagePreview"
                max-width="200"
                class="mt-3"
                contain
              ></v-img>
            </div>
          </v-col>
        </v-row>

        <div class="d-flex gap-2 mt-5">
          <v-btn
            color="primary"
            type="submit"
            :loading="loading"
            prepend-icon="mdi-content-save"
          >
            Simpan
          </v-btn>
          <v-btn
            color="error"
            @click="$router.go(-1)"
          >
            Batal
          </v-btn>
        </div>
      </v-form>
    </v-card>
  </div>
</template>

<script>
import axios from 'axios';
import Swal from 'sweetalert2';

export default {
  data() {
    return {
      formData: {
        nama_institusi: '',
        alamat: '',
        email: '',
        no_telp: '',
        logo: null
      },
      imagePreview: null,
      loading: false,
      rules: {
        required: v => !!v || 'Field ini harus diisi',
        email: v => /.+@.+\..+/.test(v) || 'Email harus valid'
      }
    }
  },

  methods: {
    previewImage() {
      if (this.formData.logo) {
        this.imagePreview = URL.createObjectURL(this.formData.logo)
      } else {
        this.imagePreview = null
      }
    },

    async submitForm() {
      try {
        if (!this.$refs.form.validate()) return

        this.loading = true

        const formData = new FormData()
        formData.append('nama_institusi', this.formData.nama_institusi)
        formData.append('alamat', this.formData.alamat)
        formData.append('email', this.formData.email)
        formData.append('no_telp', this.formData.no_telp)
        formData.append('logo', this.formData.logo)

        await axios.post('/api/create-institusi', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        await Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: 'Data institusi berhasil ditambahkan',
          showConfirmButton: false,
          timer: 1500
        })

        this.$router.push('/admin/institusi')

      } catch (error) {
        console.error('Error:', error)
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Terjadi kesalahan saat menyimpan data',
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
</style>`