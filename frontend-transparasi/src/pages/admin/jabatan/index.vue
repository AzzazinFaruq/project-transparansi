<template>
  <v-container>
  <div class="mt-5">
    <div class="d-flex justify-space-between align-center mb-5">
      <div>
        <p class="text-h5 font-weight-bold">Manajemen Jabatan</p>
      </div>
      <v-dialog v-model="dialog" max-width="500px" persistent>
        <template v-slot:activator="{ props }">
          <v-btn
            color="#387144"
            style="color: white"
            v-bind="props"
            prepend-icon="mdi-plus"
          >
            Tambah Jabatan
          </v-btn>
        </template>
        <v-card>
          <div class="d-flex justify-space-between align-center pa-5">
            <p class="text-h6 font-weight-bold mb-0">Tambah Jabatan Baru</p>
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
            <v-form ref="form" @submit.prevent="saveJabatan">
              <div class="mb-5">
                <label class="label-form mb-2 d-block">Nama Jabatan</label>
                <v-text-field
                  v-model="formData.jabatan"
                  placeholder="Masukkan nama jabatan"
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
              @click="saveJabatan"
              :loading="loading"
              prepend-icon="mdi-content-save"
            >
              Simpan
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>

    <v-divider class="mb-5"></v-divider>
    <!-- Table -->
    <v-table class="no-divider">
      <thead>
        <tr>
          <th class="text-center" style="width: 80px">No</th>
          <th class="text-left">Nama Jabatan</th>
          <th class="text-center" style="width: 100px">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in displayedItems" :key="item.Id">
          <td class="text-center">{{ startNumber + index }}</td>
          <td>{{ item.jabatan }}</td>
          <td class="text-center">
            <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon @click="deleteItem(item.Id)">
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
          </td>
        </tr>
      </tbody>
    </v-table>

    <!-- Pagination -->
  </div>
        <v-divider class="my-3"></v-divider>
        <div class="d-flex justify-end align-center">
          <p class="mr-3">Data/Halaman :</p>
          <div class="mr-5" style="width: 100px;height: 40px;">
            <v-select
            density="compact"
            class="custom-outline"
            v-model="itemsPerPage"
            :items="[5, 10, 25, 50, 100]"
            variant="outlined"
          ></v-select>
          </div>
          <p class="mr-5">{{ currentPage }} dari {{ totalPages }} Halaman</p>
          <v-pagination
            class="custom-pagination"
            :total-visible="0"
            border
            v-model="currentPage"
            :length="totalPages"
            @input="changePage"
          >
          </v-pagination>

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
      jabatan: [],
      formData: {
        jabatan: ''
      },
      rules: {
        required: v => !!v || 'Field ini harus diisi'
      },
      // Pagination
      page: 1,
      itemsPerPage: 10,
      totalItems: 0
    }
  },

  computed: {
    totalPages() {
      return Math.ceil(this.jabatan.length / this.itemsPerPage)
    },
    displayedItems() {
      const start = (this.page - 1) * this.itemsPerPage
      const end = start + this.itemsPerPage
      return this.jabatan.slice(start, end)
    },
    startNumber() {
      return (this.page - 1) * this.itemsPerPage + 1
    }
  },

  mounted() {
    this.getJabatan()
  },

  methods: {
    getJabatan() {
      this.loading = true
      axios.get('/api/index-jabatan')
        .then(res => {
          this.jabatan = res.data.data
          this.totalItems = this.jabatan.length
        })
        .catch(error => {
          console.error('Error:', error)
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal mengambil data jabatan',
          })
        })
        .finally(() => {
          this.loading = false
        })
    },

    handlePageChange(page) {
      this.page = page
    },

    saveJabatan() {
      if (!this.$refs.form.validate()) return

      this.loading = true
      axios.post('/api/create-jabatan', this.formData)
        .then(() => {
          Swal.fire({
            icon: 'success',
            title: 'Berhasil!',
            text: 'Jabatan berhasil ditambahkan',
            showConfirmButton: false,
            timer: 1500
          })
          this.closeDialog()
          this.getJabatan()
        })
        .catch(error => {
          console.error('Error:', error)
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Gagal menambahkan jabatan',
          })
        })
        .finally(() => {
          this.loading = false
        })
    },

    deleteItem(id) {
      Swal.fire({
        title: 'Hapus Jabatan?',
        text: "Data yang dihapus tidak dapat dikembalikan!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#BF3232',
        cancelButtonColor: '#808080',
        confirmButtonText: 'Ya, Hapus!',
        cancelButtonText: 'Batal',
        reverseButtons: true
      }).then((result) => {
        if (result.isConfirmed) {
          axios.delete(`/api/jabatan/${id}`)
            .then(() => {
              Swal.fire({
                icon: 'success',
                title: 'Terhapus!',
                text: 'Jabatan berhasil dihapus',
                showConfirmButton: false,
                timer: 1500
              })
              this.getJabatan()
            })
            .catch(error => {
              console.error('Error:', error)
              Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'Gagal menghapus jabatan',
              })
            })
        }
      })
    },

    closeDialog() {
      this.dialog = false
      this.$refs.form.reset()
      this.formData = {
        jabatan: ''
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


</style>