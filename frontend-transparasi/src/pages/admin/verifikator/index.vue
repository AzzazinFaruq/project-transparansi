<template>
  <v-container fluid class="">
    <div class="d-flex justify-space-between align-center mb-5 px-4">
      <div>
        <p class="text-h5 font-weight-bold">Manajemen Dinas Verifikator</p>
        <p class="text-subtitle-1 text-grey">Kelola data dinas verifikator</p>
      </div>
      <v-dialog v-model="dialog" max-width="500px" persistent>
        <template v-slot:activator="{ props }">
          <v-btn
            color="#387144"
            style="color: white"
            v-bind="props"
            prepend-icon="mdi-plus"
          >
            Tambah Dinas Verifikator
          </v-btn>
        </template>
        <v-card>
          <div class="d-flex justify-space-between align-center pa-5">
            <p class="text-h6 font-weight-bold mb-0">Tambah Dinas Verifikator Baru</p>
            <v-btn
              icon
              variant="text"
              @click="closeDialog"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </div>
          <v-divider></v-divider>
          <v-card-text class="pa-5">
            <v-form ref="form" @submit.prevent="saveDinasVerifikator">
              <div class="mb-5">
                <label class="label-form mb-2 d-block">Nama Dinas Verifikator</label>
                <v-text-field
                  v-model="formData.nama_dinas_verifikator"
                  placeholder="Masukkan nama Dinas Verifikator"
                  variant="outlined"
                  density="compact"
                  :rules="[rules.required]"
                ></v-text-field>
              </div>
            </v-form>
          </v-card-text>
          <v-card-actions class="pa-5 pt-0">
            <v-spacer></v-spacer>
            <v-btn
              color="#BF3232"
              style="color: white"
              @click="closeDialog"
            >
              Batal
            </v-btn>
            <v-btn
              color="#387144"
              style="color: white"
              @click="saveDinasVerifikator"
              :loading="loading"
              prepend-icon="mdi-content-save"
            >
              Simpan
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>

    <div class="table-responsive">
      <v-table class="no-divider">
        <thead>
          <tr>
            <th class="text-center" style="width: 80px">No</th>
            <th class="text-left">Nama Dinas</th>
            <th class="text-center" style="width: 100px">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in items" :key="item.Id">
            <td class="text-center">{{ index + 1 }}</td>
            <td>{{ item.nama_dinas_verifikator }}</td>
            <td class="text-center">
              <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon @click="deleteItem(item.Id)">
                <v-icon>mdi-delete</v-icon>
              </v-btn>
            </td>
          </tr>
        </tbody>
      </v-table>
    </div>
  </v-container>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2'

export default {
  data() {
    return {
      dialog: false,
      loading: false,
      items: [],
      formData: {
        nama_dinas_verifikator: ''
      },
      rules: {
        required: value => !!value || 'Wajib diisi'
      }
    }
  },

  mounted() {
    this.getItems()
  },

  methods: {
    async getItems() {
      try {
        const response = await axios.get('http://localhost:8000/api/index-dinas-verifikator')
        this.items = response.data.data
      } catch (error) {
        console.error('Error:', error)
      }
    },

    async saveDinasVerifikator() {
      if (!this.$refs.form.validate()) return

      this.loading = true
      try {
        await axios.post('http://localhost:8000/api/create-dinas-verifikator', {
          nama_dinas_verifikator: this.formData.nama_dinas_verifikator
        })

        this.closeDialog()
        this.getItems()

        Swal.fire({
          icon: 'success',
          title: 'Berhasil',
          text: 'Data dinas verifikator berhasil ditambahkan',
          showConfirmButton: false,
          timer: 1500
        })
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
    },

    deleteItem(id) {
      Swal.fire({
        title: 'Apakah anda yakin?',
        text: "Data yang dihapus tidak dapat dikembalikan!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#387144',
        cancelButtonColor: '#BF3232',
        confirmButtonText: 'Ya, hapus!',
        cancelButtonText: 'Batal'
      }).then((result) => {
        if (result.isConfirmed) {
          axios.delete(`http://localhost:8000/api/dinas-verifikator/${id}`)
            .then(() => {
              this.getItems()
              Swal.fire({
                icon: 'success',
                title: 'Berhasil',
                text: 'Data dinas verifikator berhasil dihapus',
                showConfirmButton: false,
                timer: 1500
              })
            })
            .catch(error => {
              console.error('Error:', error)
              Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'Gagal menghapus data',
              })
            })
        }
      })
    },

    closeDialog() {
      this.dialog = false
      this.$refs.form.reset()
      this.formData = {
        nama_dinas_verifikator: ''
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

.v-btn {
  text-transform: none !important;
}

.table-responsive {
  width: 100%;
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
}

@media (max-width: 600px) {
  .d-flex.justify-space-between {
    flex-direction: column;
    gap: 1rem;
  }

  .d-flex.justify-end {
    flex-direction: column;
    align-items: stretch;
    gap: 1rem;
  }

  .mr-3, .mr-5 {
    margin-right: 0 !important;
    margin-bottom: 0.5rem;
  }
}

.table-responsive::-webkit-scrollbar {
  height: 6px;
}

.table-responsive::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.table-responsive::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 3px;
}

.table-responsive::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>
