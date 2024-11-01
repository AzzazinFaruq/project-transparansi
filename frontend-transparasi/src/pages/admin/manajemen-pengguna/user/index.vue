<template>
  <v-container>
  <div class="mt-5">
        <div class="d-flex align-center justify-space-between mr-3 mt-2">
          <v-card-title><b>Daftar Program</b></v-card-title>
          <a href="">
            <v-icon class="">mdi-dots-vertical</v-icon>
          </a>
        </div>
        <v-divider class="mx-2"></v-divider>
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
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon @click="deleteUser(item.Id)">
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
    }
  },
  mounted() {
    this.user();
  },
  methods: {
    deleteUser(id){
      axios.delete(`api/user/deleteuser/${id}`)
      .then(res=>{
        this.$router.go(0)
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
}
</script>
