<template>
  <div>
    <v-card class="pa-5">
      <div class="d-flex justify-space-between align-center mb-5">
        <div>
          <p class="text-h5 font-weight-bold">Edit Pengguna</p>
          <p class="text-subtitle-1 text-grey">Edit informasi pengguna</p>
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
          <v-form>
            <!-- Informasi Dasar -->
            <div class="mb-5">
              <label class="label-form mb-2 d-block">Nama</label>
              <v-text-field
                v-model="user.username"
                placeholder="Masukkan nama"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Email</label>
              <v-text-field
                v-model="user.email"
                placeholder="Masukkan email"
                variant="outlined"
                density="compact"
                type="email"
              ></v-text-field>
            </div>

            <!-- Detail Profil -->
            <h2 class="text-h6 font-weight-bold mb-3">Detail Profil</h2>
            <v-divider class="mb-5"></v-divider>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">Alamat</label>
              <v-textarea
                v-model="user.alamat"
                placeholder="Masukkan alamat"
                variant="outlined"
                density="compact"
                rows="3"
              ></v-textarea>
            </div>

            <div class="mb-5">
              <label class="label-form mb-2 d-block">No. HP</label>
              <v-text-field
                v-model="user.no_hp"
                placeholder="Masukkan nomor HP"
                variant="outlined"
                density="compact"
              ></v-text-field>
            </div>

            <!-- Tombol Aksi -->
            <div class="d-flex gap-2 mt-5">
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
import Swal from 'sweetalert2';

export default {
  data() {
    return {
      user: {
        username: '',
        email: '',
        password: '',
        no_hp: '',
        alamat: ''
      }
    }
  },

  mounted() {
    this.getUser();
  },

  methods: {
    getUser() {
      axios.get(`/api/user/${this.$route.params.id}`)
        .then(res => {
          this.user = res.data.data;
        })
        .catch(error => {
          console.error('Error:', error);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal mengambil data pengguna',
          });
        });
    },

    update() {
      axios.put(`/api/user/edituser/${this.$route.params.id}`, this.user)
        .then(() => {
          Swal.fire({
            icon: 'success',
            title: 'Berhasil!',
            text: 'Data pengguna berhasil diperbarui',
            showConfirmButton: false,
            timer: 1500
          }).then(() => {
            this.$router.go(-1);
          });
        })
        .catch(error => {
          console.error('Error:', error);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal memperbarui data pengguna',
          });
        });
    },

    back() {
      this.$router.go(-1);
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
</style>
