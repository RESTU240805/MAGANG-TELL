<template>
  <div class="min-h-screen bg-gray-100 flex">
    <aside class="w-64 bg-gray-950 text-white flex flex-col fixed h-full z-40">
      <div class="p-6 border-b border-gray-800">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-green-600 rounded-xl flex items-center justify-center font-black text-sm">T</div>
          <div>
            <p class="font-black text-sm">TELPP</p>
            <p class="text-xs text-gray-500">Management Profile</p>
          </div>
        </div>
      </div>
      <nav class="flex-1 p-4 space-y-1">
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-2 px-2">OVERVIEW</p>
        <RouterLink to="/admin/dashboard"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📊 Dashboard
        </RouterLink>
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-4 px-2">CONTENT ENGINE</p>
        <RouterLink to="/admin/news"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📰 Corporate News
        </RouterLink>
        <RouterLink to="/admin/slider"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🖼️ Product Slider
        </RouterLink>
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📝 Product Page
        </RouterLink>
        <RouterLink to="/admin/about"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏢 About Section
        </RouterLink>
        <RouterLink to="/admin/team-members"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          👥 Our Team
        </RouterLink>
        <RouterLink to="/admin/our-company"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏛️ Our Company
        </RouterLink>
        <RouterLink to="/admin/menus"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          📋 Menu
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <main class="flex-1 ml-64 p-10">
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Menu</h1>
        </div>
        <button @click="openCreate"
          class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-2">
          <span class="text-lg leading-none">+</span> Tambah Menu
        </button>
      </div>

      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <div v-if="loading" class="flex items-center justify-center py-20 gap-3 text-gray-400">
        <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
        Memuat data...
      </div>

      <div v-else-if="menus.length === 0"
        class="bg-white rounded-2xl p-16 text-center border border-gray-100">
        <p class="text-4xl mb-4">📋</p>
        <p class="text-gray-500 mb-4">Belum ada menu.</p>
        <button @click="openCreate"
          class="bg-green-600 text-white px-6 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Menu
        </button>
      </div>

      <div v-else class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <table class="w-full text-sm">
          <thead class="bg-gray-50 border-b border-gray-100">
            <tr>
              <th class="text-left px-6 py-3 font-semibold text-gray-600">Nama</th>
              <th class="text-left px-6 py-3 font-semibold text-gray-600">URL</th>
              <th class="text-center px-6 py-3 font-semibold text-gray-600">Status</th>
              <th class="text-center px-6 py-3 font-semibold text-gray-600">Urutan</th>
              <th class="text-center px-6 py-3 font-semibold text-gray-600">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-50">
            <tr v-for="menu in flatMenus" :key="menu.id"
              :class="menu.parent_id ? 'bg-gray-50/50' : ''">
              <td class="px-6 py-3">
                <span :class="menu.parent_id ? 'ml-6 text-gray-500' : 'font-semibold text-gray-900'">
                  {{ menu.parent_id ? '↳ ' : '' }}{{ menu.name }}
                </span>
              </td>
              <td class="px-6 py-3 text-gray-500">{{ menu.url }}</td>
              <td class="px-6 py-3 text-center">
                <button @click="toggleActive(menu)"
                  :class="[menu.is_active ? 'bg-green-500 text-white' : 'bg-gray-200 text-gray-500', 'relative inline-flex h-6 w-11 items-center rounded-full transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-offset-2']"
                  :disabled="togglingId === menu.id">
                  <span :class="[menu.is_active ? 'translate-x-6' : 'translate-x-1', 'inline-block h-4 w-4 rounded-full bg-white shadow transition-transform duration-200']"></span>
                </button>
              </td>
              <td class="px-6 py-3 text-center text-gray-500">{{ menu.sort_order }}</td>
              <td class="px-6 py-3 text-center">
                <div class="flex gap-2 justify-center">
                  <button @click="openEdit(menu)"
                    class="bg-blue-50 text-blue-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-blue-100 transition">Edit</button>
                  <button @click="confirmDelete(menu)"
                    class="bg-red-50 text-red-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-red-100 transition">Hapus</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>

    <!-- Modal Form -->
    <Transition name="modal">
      <div v-if="showModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="closeModal">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-lg max-h-[90vh] overflow-y-auto">
          <div class="flex items-center justify-between p-6 border-b border-gray-100">
            <h2 class="text-xl font-black text-gray-900">
              {{ isEdit ? '✏️ Edit Menu' : '➕ Tambah Menu' }}
            </h2>
            <button @click="closeModal" class="text-gray-400 hover:text-gray-600 text-2xl leading-none">×</button>
          </div>
          <form @submit.prevent="submitForm" class="p-6 space-y-5">
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Nama Menu <span class="text-red-500">*</span></label>
              <input v-model="form.name" type="text" placeholder="Nama menu..." required
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">URL <span class="text-red-500">*</span></label>
              <input v-model="form.url" type="text" placeholder="/page-path" required
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Parent Menu</label>
              <select v-model="form.parent_id"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400">
                <option :value="null">— Tidak ada (menu utama) —</option>
                <option v-for="p in parentOptions" :key="p.id" :value="p.id">{{ p.name }}</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Urutan</label>
              <input v-model.number="form.sort_order" type="number" min="0"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
            </div>
            <div class="flex items-center gap-3">
              <input id="is_active" v-model="form.is_active" type="checkbox"
                class="w-4 h-4 rounded border-gray-300 text-green-600 focus:ring-green-400" />
              <label for="is_active" class="text-sm font-semibold text-gray-700">Aktif</label>
            </div>
            <div v-if="formError" class="bg-red-50 text-red-600 text-sm px-4 py-3 rounded-xl border border-red-200">
              ❌ {{ formError }}
            </div>
            <div class="flex gap-3 pt-2">
              <button type="button" @click="closeModal"
                class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">Batal</button>
              <button type="submit" :disabled="submitting"
                class="flex-1 bg-green-600 text-white py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition disabled:opacity-60 flex items-center justify-center gap-2">
                <div v-if="submitting" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                {{ submitting ? 'Menyimpan...' : (isEdit ? 'Simpan Perubahan' : 'Tambah Menu') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Transition>

    <!-- Modal Hapus -->
    <Transition name="modal">
      <div v-if="showDeleteModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="showDeleteModal = false">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-sm p-6 text-center">
          <p class="text-4xl mb-3">🗑️</p>
          <h3 class="text-lg font-black text-gray-900 mb-1">Hapus Menu?</h3>
          <p class="text-sm text-gray-500 mb-6">Menu akan dihapus permanen.</p>
          <div class="flex gap-3">
            <button @click="showDeleteModal = false"
              class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">Batal</button>
            <button @click="doDelete" :disabled="submitting"
              class="flex-1 bg-red-600 text-white py-2.5 rounded-xl text-sm font-semibold hover:bg-red-700 transition disabled:opacity-60">
              {{ submitting ? 'Menghapus...' : 'Hapus' }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const router = useRouter()
const menus = ref([])
const loading = ref(true)
const submitting = ref(false)
const showModal = ref(false)
const showDeleteModal = ref(false)
const isEdit = ref(false)
const selectedMenu = ref(null)
const formError = ref('')
const alert = ref({ show: false, type: 'success', message: '' })

const form = ref({ name: '', url: '', parent_id: null, sort_order: 0, is_active: true })
const togglingId = ref(null)

const flatMenus = computed(() => {
  const result = []
  const flatten = (items, depth = 0) => {
    for (const item of items) {
      result.push({ ...item, depth })
      if (item.children && item.children.length > 0) {
        flatten(item.children, depth + 1)
      }
    }
  }
  flatten(menus.value)
  return result
})

const parentOptions = computed(() => {
  const result = []
  const collect = (items) => {
    for (const item of items) {
      if (item.id !== selectedMenu.value?.id) {
        result.push({ id: item.id, name: item.name })
        if (item.children && item.children.length > 0) {
          collect(item.children)
        }
      }
    }
  }
  collect(menus.value)
  return result
})

const showAlert = (type, message) => {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

const fetchMenus = async () => {
  loading.value = true
  try {
    const res = await api.get('/admin/menus')
    menus.value = res.data?.data || []
  } catch { showAlert('error', 'Gagal memuat menu') }
  finally { loading.value = false }
}

const toggleActive = async (menu) => {
  togglingId.value = menu.id
  try {
    await api.put(`/admin/menus/${menu.id}`, {
      name: menu.name,
      url: menu.url,
      parent_id: menu.parent_id,
      sort_order: menu.sort_order,
      is_active: !menu.is_active,
    })
    showAlert('success', `"${menu.name}" ${!menu.is_active ? 'diaktifkan' : 'dinonaktifkan'}`)
    await fetchMenus()
  } catch {
    showAlert('error', 'Gagal mengubah status')
  } finally { togglingId.value = null }
}

const openCreate = () => {
  isEdit.value = false
  form.value = { name: '', url: '', parent_id: null, sort_order: 0, is_active: true }
  formError.value = ''
  showModal.value = true
}

const openEdit = (menu) => {
  isEdit.value = true
  selectedMenu.value = menu
  form.value = {
    name: menu.name || '',
    url: menu.url || '',
    parent_id: menu.parent_id || null,
    sort_order: menu.sort_order || 0,
    is_active: menu.is_active !== undefined ? menu.is_active : true,
  }
  formError.value = ''
  showModal.value = true
}

const closeModal = () => { showModal.value = false; formError.value = '' }

const submitForm = async () => {
  if (!form.value.name.trim()) {
    formError.value = 'Nama menu wajib diisi'; return
  }
  if (!form.value.url.trim()) {
    formError.value = 'URL wajib diisi'; return
  }
  submitting.value = true; formError.value = ''
  try {
    const payload = {
      name: form.value.name,
      url: form.value.url,
      parent_id: form.value.parent_id,
      sort_order: form.value.sort_order,
      is_active: form.value.is_active,
    }
    if (isEdit.value) {
      await api.put(`/admin/menus/${selectedMenu.value.id}`, payload)
      showAlert('success', 'Menu berhasil diperbarui!')
    } else {
      await api.post('/admin/menus', payload)
      showAlert('success', 'Menu berhasil ditambahkan!')
    }
    closeModal(); await fetchMenus()
  } catch (e) {
    console.error('Submit error:', e)
    formError.value = e.response?.data?.error || e.response?.data?.message || e.message || 'Terjadi kesalahan. Coba lagi.'
  } finally { submitting.value = false }
}

const confirmDelete = (menu) => { selectedMenu.value = menu; showDeleteModal.value = true }
const doDelete = async () => {
  submitting.value = true
  try {
    await api.delete(`/admin/menus/${selectedMenu.value.id}`)
    showDeleteModal.value = false
    showAlert('success', 'Menu berhasil dihapus!')
    await fetchMenus()
  } catch { showAlert('error', 'Gagal menghapus menu') }
  finally { submitting.value = false }
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}

onMounted(() => {
  if (!localStorage.getItem('token')) { router.push('/admin/login'); return }
  fetchMenus()
})
</script>

<style scoped>
.modal-enter-active, .modal-leave-active { transition: opacity 0.2s ease; }
.modal-enter-from, .modal-leave-to { opacity: 0; }
</style>
