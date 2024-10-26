<template>
  <div class="mt-3">
        <div class="d-flex align-center justify-space-between mr-3 mt-2">
          <v-card-title><b>Daftar Program</b></v-card-title>
          <a href="">
            <v-icon class="">mdi-dots-vertical</v-icon>
          </a>
        </div>
        <v-divider class="mx-2"></v-divider>
        <div class="">
          <v-table class="no-divider ">
            <thead class="">
              <tr >
                <th class=" font-weight-bold">
                  Tanggal
                </th>
                <th class=" font-weight-bold">
                  Nama
                </th>
                <th class=" font-weight-bold">
                  Status
                </th>
                <th class=" font-weight-bold">
                  Action
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="item in Userlist"
                :key="item.name"
                >
              <td>{{ item.CreatedAt}}</td>
              <td>{{ item.nama_program }}</td>
              <td>{{ item.status }}</td>
              <td>
                <v-btn class="rounded-lg" style="background-color:#3884B0;color: white;text-transform: none;" >
                  Detail
                </v-btn>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #387144;color: white;" icon>
                  <v-icon>mdi-check-bold</v-icon>
                </v-btn>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon>
                  <b style="font-weight: 900;">X</b>
                </v-btn>
              </td>

              </tr>
              </tbody>
          </v-table>
        </div>
    </div>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      headers: [
        { title: "Nama", value: "username" },
        { title: "Jabatan", value: "" },
        { title: "Email", value: "email" },
        { title: "No Hp", value: "no_hp" },
        { title: "Action", value: "actions" },
      ],
      Userlist:[],
    }
  },
  mounted() {
    this.user();
  },
  methods: {
    user(){
      axios.get("/api/index-program")
      .then(res=>{
        console.log(res.data.data)
        this.Userlist=res.data.data
      })
    }
  },
}
</script>
