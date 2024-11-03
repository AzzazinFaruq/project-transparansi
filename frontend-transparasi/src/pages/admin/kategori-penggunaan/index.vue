<template>
  <v-container fluid class="">
    <div class="d-flex justify-space-between align-center mb-5 px-4">
      <div>
        <p class="text-h5 font-weight-bold">Kategori Penggunaan</p>
        <p class="text-subtitle-1 text-grey">Kelola Kategori Penggunaan</p>
      </div>
      <v-dialog v-model="dialog" max-width="500px" persistent>
        <template v-slot:activator="{ props }">
          <v-btn
            color="#387144"
            style="color: white"
            v-bind="props"
            prepend-icon="mdi-plus"
          >
            Tambah Kategori
          </v-btn>
        </template>
        <v-card>
          <div class="d-flex justify-space-between align-center pa-5">
            <p class="text-h6 font-weight-bold mb-0">Tambah Kategori Baru</p>
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
                <label class="label-form mb-2 d-block">Nama Kategori</label>
                <v-text-field
                  v-model="formData.kategori  "
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
              color="red"
              style="color: white"
              @click="closeDialog"
            >
              Batal
            </v-btn>
            <v-btn
              color="green"
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

    <div class="table-responsive">
      <v-table class="no-divider">
        <thead>
          <tr>
            <th class="text-center" style="width: 80px">No</th>
            <th class="text-left">Kategori</th>
            <th class="text-center" style="width: 100px">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in paginatedItems" :key="item.Id">
            <td class="text-center">{{ index +1 }}</td>
            <td>{{ item.kategori }}</td>
            <td class="text-center">
              <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon @click="deleteItem(item.Id)">
                    <v-icon>mdi-delete</v-icon>
                  </v-btn>
            </td>
          </tr>
        </tbody>
      </v-table>
    </div>

    <v-divider class="my-3"></v-divider>
    <div class="d-flex flex-wrap justify-end align-center">
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
        :total-visible="$vuetify.display.mdAndUp ? 7 : 3"
        border
        v-model="currentPage"
        :length="totalPages"
        @input="handlePageChange"
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
        kategori: ''
      },
      rules: {
        required: v => !!v || 'Field ini harus diisi'
      },
      // Pagination
      currentPage: 1,
      itemsPerPage: 10,
    }
  },

  computed: {
    totalPages() {
      return Math.ceil(this.jabatan.length / this.itemsPerPage);
    },
    paginatedItems() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.jabatan.slice(start, end);
    },
  },

  mounted() {
    this.getJabatan()
  },

  methods: {
    getJabatan() {
      this.loading = true
      axios.get('/api/index-kategori-penggunaan')
        .then(res => {
          this.jabatan = res.data.data
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
      axios.post('/api/kategori-penggunaan', this.formData)
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
          axios.delete(`/api/kategori-penggunaan/${id}`)
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

/* Responsive styles */
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

  .custom-pagination {
    width: 100%;
    justify-content: center;
  }
}

/* Custom scrollbar for better UX */
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