<template>
  <div class="">
  <div class="mt-3">
    <div class="d-flex align-center justify-space-between mr-3 mt-2">
      <v-card-title><b>Daftar Institusi</b></v-card-title>
      <v-btn color="#387144" prepend-icon="mdi-plus" style="text-transform: none;" href="/admin/institusi/tambah">Institusi</v-btn>
    </div>
  </div>
  <VDivider/>
  <div class="">
          <v-table class="no-divider ">
            <thead style="">
              <tr class="">
                <th class=" font-weight-bold">
                  Tanggal
                </th>
                <th class=" font-weight-bold">
                  Nama Institusi
                </th>
                <th class=" font-weight-bold">
                  Action
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="(item, index) in paginatedItems" :key="index"
                >
              <td>{{ item.created_at}}</td>
              <td>{{ item.nama_institusi }}</td>
              <td>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #3884B0;color: white;" icon @click="edit(item.id)">
                  <v-icon>mdi-pencil</v-icon>
                </v-btn>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon @click="deleteConfirmation(item.id)">
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
              </td>

              </tr>
              </tbody>
          </v-table>
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
      </div>
</template>
<script>
import axios from 'axios';
import Swal from 'sweetalert2';
export default {
  data(){
    return{
      headers:[],
      institusi:[],
      currentPage:1,
      itemsPerPage:10,
    }
  },
  mounted(){
    this.getInstitusi();
  },
  methods:{
    changePage(page) {
      this.currentPage = page;
    },
    getInstitusi(){
      axios.get('/api/institusi')
      .then((res)=>{
        this.institusi = res.data.data;
      })
      .catch((err)=>{
        console.log(err);
      })
    },
    deleteConfirmation(id) {
      Swal.fire({
        title: 'Hapus Data?',
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
          // Tampilkan loading
          Swal.fire({
            title: 'Menghapus Data...',
            allowOutsideClick: false,
            didOpen: () => {
              Swal.showLoading()
            }
          })

          // Proses delete
          axios.delete(`/api/institusi/${id}`)
            .then(() => {
              Swal.fire({
                icon: 'success',
                title: 'Terhapus!',
                text: 'Data berhasil dihapus',
                showConfirmButton: false,
                timer: 1500
              }).then(() => {
                // Refresh data
                this.$router.go(0)
              })
            })
            .catch((error) => {
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
    edit(id){
      this.$router.push(`/admin/institusi/edit/${id}`)
    }
  },
  computed: {
    totalPages() {
      return Math.ceil(this.institusi.length / this.itemsPerPage);
    },
    paginatedItems() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.institusi.slice(start, end);
    },
  },
}
</script>