<template>
  <v-container>
  <div class="mt-5">
        <div class="d-flex align-center justify-start mr-3 mt-2">
          <h2>Daftar User</h2>
        </div>
        <div class="d-flex justify-space-between">
          <div class=""> 
            <v-text-field density="compact" v-model="search" prepend-inner-icon="mdi-magnify" label="Cari User" variant="outlined" style="width: 300px;" @keyup.enter="searchUser"></v-text-field>
          </div>
          </div>
          <v-divider></v-divider>
          <div class="">
          <v-table class="no-divider ">
            <thead style="">
              <tr class="">
                <th class=" font-weight-bold">
                  Nama
                </th>
                <th class=" font-weight-bold">
                  Email
                </th>
                <th class=" font-weight-bold">
                  Aksi
                </th>
              </tr>
            </thead>

              <tbody>

                <tr
                v-for="(item, index) in paginatedItems" :key="index"
                >

              <td>{{ item.username}}</td>
              <td>{{ item.email }}</td>
              <td>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #3884B0;color: white;" icon @click="edit(item.Id)">
                  <v-icon>mdi-pencil</v-icon>
                </v-btn>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon @click="deleteConfirmation(item.Id)">
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
              </td>

              </tr>
              </tbody>
          </v-table>
          <div class="d-flex justify-center align-center" style="height: 100px;" v-if="notFound == true">
                  <p>Tidak Ada Data Ditemukan</p>
                </div>
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
  </v-container>
</template>
<script>
import axios from 'axios';
import Swal from 'sweetalert2'

export default {
  data() {
    return {
      notFound:false,
      headers: [
        { title: "Nama", value: "username" },
        { title: "Jabatan", value: "" },
        { title: "Email", value: "email" },
        { title: "No Hp", value: "no_hp" },
        { title: "Action", value: "actions" },
      ],
      Userlist:[],
      currentPage: 1,
      itemsPerPage: 10,
      search:'',
    }
  },
  mounted() {
    this.user();
  },
  methods: {
    searchUser(){
      axios.get(`/api/user/username?username=${this.search}&role=3`)
      .then(res=>{
        this.Userlist = res.data.data
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
          axios.delete(`/api/user/deleteuser/${id}`)
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
    user(){
      axios.get("/api/user/byrole/3")
      .then(res=>{
        console.log(res.data.data)
        this.Userlist=res.data.data
      })
      .catch(err=>{
        this.notFound=true
      })
    },
    changePage(page) {
      this.currentPage = page;
    },
    edit(id){
      this.$router.push(`/admin/manajemen-pengguna/user/${id}`)
    }
  },
  computed: {
    totalPages() {
      return Math.ceil(this.Userlist.length / this.itemsPerPage);
    },
    paginatedItems() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.Userlist.slice(start, end);
    },
  },
  watch: {
    search(newVal) {
      if (!newVal || newVal.trim() === '') {
        this.user();
        this.currentPage = 1; // Reset ke halaman pertama
      }
    }
  }
}
</script>
