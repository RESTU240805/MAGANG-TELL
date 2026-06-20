<template>
  <nav class="fixed top-0 left-0 right-0 w-full z-50 shadow-sm bg-white">
    <div class="w-full px-4 sm:px-10 py-3 flex items-center justify-between">

      <!-- Logo -->
      <RouterLink to="/" class="flex items-center gap-3 sm:gap-4">
        <img src="/images/logo-telpp.png" alt="TeL Logo" class="h-10 sm:h-12 w-auto" />
        <span class="text-[10px] sm:text-[11px] font-bold text-gray-800 tracking-[1.5px] sm:tracking-[2px] uppercase leading-snug">
          PT Tanjungenim Lestari<br />Pulp and Paper
        </span>
      </RouterLink>

      <!-- Hamburger button (mobile) -->
      <button
        class="hamburger-btn"
        :class="{ 'is-open': mobileOpen }"
        @click="toggleMobile"
        aria-label="Toggle menu"
      >
        <span></span>
        <span></span>
        <span></span>
      </button>

      <!-- Desktop Menu -->
      <div class="desktop-menu">
        <template v-for="item in menus" :key="item.id">
          <!-- Dropdown (has children) -->
          <div v-if="item.children?.length" class="nav-dropdown-wrapper">
            <button
              class="nav-link nav-dropdown-trigger"
              :class="{ 'nav-link--active': isChildActive(item) }"
            >
              {{ item.name }}
              <span style="font-size:9px; margin-top:1px;">▾</span>
            </button>
            <div class="nav-dropdown-menu">
              <template v-for="child in item.children" :key="child.id">
                <!-- Nested dropdown -->
                <div v-if="child.children?.length" class="subdropdown-wrapper">
                  <button class="dropdown-item subdropdown-trigger w-full text-left border-none bg-transparent cursor-pointer flex justify-between items-center">
                    {{ child.name }}
                    <span style="color:#aaa; margin-left:8px;">›</span>
                  </button>
                  <div class="subdropdown-menu">
                    <RouterLink
                      v-for="gc in child.children" :key="gc.id"
                      :to="gc.url"
                      class="dropdown-item"
                      :class="{ 'dropdown-item--active': isActive(gc.url) }"
                    >
                      {{ gc.name }}
                    </RouterLink>
                  </div>
                </div>
                <!-- External child -->
                <a
                  v-else-if="isExternal(child.url)"
                  :href="child.url"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="dropdown-item"
                >
                  {{ child.name }}
                </a>
                <!-- Internal child -->
                <RouterLink
                  v-else
                  :to="child.url"
                  class="dropdown-item"
                  :class="{ 'dropdown-item--active': isActive(child.url) }"
                >
                  {{ child.name }}
                </RouterLink>
              </template>
            </div>
          </div>

          <!-- External top link -->
          <a
            v-else-if="isExternal(item.url)"
            :href="item.url"
            target="_blank"
            rel="noopener noreferrer"
            class="nav-link"
          >
            {{ item.name }}
          </a>

          <!-- Internal top link -->
          <RouterLink
            v-else
            :to="item.url"
            class="nav-link"
            :class="{ 'nav-link--active': isActive(item.url) }"
          >
            {{ item.name }}
          </RouterLink>
        </template>
      </div>
    </div>

    <!-- Mobile Menu Overlay -->
    <Transition name="fade">
      <div
        v-if="mobileOpen"
        class="mobile-overlay"
        @click="mobileOpen = false"
      ></div>
    </Transition>

    <!-- Mobile Menu Sidebar -->
    <Transition name="slide">
      <div v-if="mobileOpen" class="mobile-sidebar">
        <div class="mobile-sidebar-header">
          <RouterLink to="/" class="flex items-center gap-3" @click="mobileOpen = false">
            <img src="/images/logo-telpp.png" alt="TeL Logo" class="h-9 w-auto" />
            <span class="text-[9px] font-bold text-gray-800 tracking-[1.5px] uppercase leading-snug">
              PT TELPP
            </span>
          </RouterLink>
          <button class="mobile-close-btn" @click="mobileOpen = false" aria-label="Close menu">
            <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
            </svg>
          </button>
        </div>

        <div class="mobile-nav-scroll">
          <template v-for="item in menus" :key="item.id">
            <!-- Mobile dropdown -->
            <template v-if="item.children?.length">
              <button
                class="mobile-link mobile-link--toggle"
                :class="{ 'mobile-link--active': isChildActive(item) }"
                @click="toggleMobileItem(item.id)"
              >
                {{ item.name }}
                <span class="mobile-toggle-icon" :class="{ 'is-open': mobileOpenItems.has(item.id) }">▾</span>
              </button>
              <Transition name="expand">
                <div v-if="mobileOpenItems.has(item.id)" class="mobile-submenu">
                  <template v-for="child in item.children" :key="child.id">
                    <!-- Nested mobile dropdown -->
                    <template v-if="child.children?.length">
                      <button
                        class="mobile-sublink mobile-link--toggle"
                        @click="toggleMobileItem(child.id)"
                      >
                        {{ child.name }}
                        <span class="mobile-toggle-icon" :class="{ 'is-open': mobileOpenItems.has(child.id) }">›</span>
                      </button>
                      <Transition name="expand">
                        <div v-if="mobileOpenItems.has(child.id)" class="mobile-submenu mobile-submenu--inner">
                          <RouterLink
                            v-for="gc in child.children" :key="gc.id"
                            :to="gc.url"
                            class="mobile-sublink"
                            @click="mobileOpen = false"
                          >
                            {{ gc.name }}
                          </RouterLink>
                        </div>
                      </Transition>
                    </template>
                    <!-- External mobile child -->
                    <a
                      v-else-if="isExternal(child.url)"
                      :href="child.url"
                      target="_blank"
                      rel="noopener noreferrer"
                      class="mobile-sublink"
                      @click="mobileOpen = false"
                    >
                      {{ child.name }}
                    </a>
                    <!-- Internal mobile child -->
                    <RouterLink
                      v-else
                      :to="child.url"
                      class="mobile-sublink"
                      @click="mobileOpen = false"
                    >
                      {{ child.name }}
                    </RouterLink>
                  </template>
                </div>
              </Transition>
            </template>

            <!-- External mobile top link -->
            <a
              v-else-if="isExternal(item.url)"
              :href="item.url"
              target="_blank"
              rel="noopener noreferrer"
              class="mobile-link"
              @click="mobileOpen = false"
            >
              {{ item.name }}
            </a>

            <!-- Internal mobile top link -->
            <RouterLink
              v-else
              :to="item.url"
              class="mobile-link"
              :class="{ 'mobile-link--active': isActive(item.url) }"
              @click="mobileOpen = false"
            >
              {{ item.name }}
            </RouterLink>
          </template>
        </div>
      </div>
    </Transition>
  </nav>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import api from '../services/api'

const route = useRoute()
const mobileOpen = ref(false)
const menus = ref([])
const mobileOpenItems = ref(new Set())

const isExternal = (url) => url && url.startsWith('http')

const isActive = (path) => {
  if (!path) { return false }
  if (path === '/') { return route.path === '/' }
  return route.path.startsWith(path)
}

const isChildActive = (item) => {
  if (!item.children) { return isActive(item.url) }
  const walk = (items) => {
    for (const c of items) {
      if (isActive(c.url)) { return true }
      if (c.children && walk(c.children)) { return true }
    }
    return false
  }
  return walk(item.children)
}

const toggleMobile = () => {
  mobileOpen.value = !mobileOpen.value
}

const toggleMobileItem = (id) => {
  const next = new Set(mobileOpenItems.value)
  if (next.has(id)) {
    next.delete(id)
  } else {
    next.add(id)
  }
  mobileOpenItems.value = next
}

const fetchMenus = async () => {
  try {
    const res = await api.get('/menus')
    menus.value = res.data?.data || []
  } catch {
    menus.value = []
  }
}

watch(() => route.path, () => {
  mobileOpen.value = false
  mobileOpenItems.value = new Set()
})

onMounted(fetchMenus)
</script>

<style scoped>
/* ── Desktop menu ── */
.desktop-menu {
  display: flex;
  align-items: stretch;
}

/* ── Hamburger button ── */
.hamburger-btn {
  display: none;
  flex-direction: column;
  justify-content: center;
  gap: 5px;
  width: 36px;
  height: 36px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 6px;
  z-index: 1001;
}

.hamburger-btn span {
  display: block;
  width: 100%;
  height: 2.5px;
  background: #333;
  border-radius: 2px;
  transition: transform 0.3s, opacity 0.3s;
}

.hamburger-btn.is-open span:nth-child(1) {
  transform: translateY(7.5px) rotate(45deg);
}

.hamburger-btn.is-open span:nth-child(2) {
  opacity: 0;
}

.hamburger-btn.is-open span:nth-child(3) {
  transform: translateY(-7.5px) rotate(-45deg);
}

/* ── Nav link ── */
.nav-link {
  color: #333333;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 1.5px;
  text-transform: uppercase;
  padding: 14px 16px;
  white-space: nowrap;
  display: flex;
  align-items: center;
  gap: 4px;
  text-decoration: none;
  border-bottom: 3px solid transparent;
  transition: color 0.2s, border-color 0.2s;
}

.nav-link:hover {
  color: #5F9E42;
}

.nav-link--active {
  color: #5F9E42;
  border-bottom: 3px solid #5F9E42;
}

/* ── Generic dropdown (replaces hardcoded group names) ── */
.nav-dropdown-wrapper {
  position: relative;
}

.nav-dropdown-trigger {
  background: none;
  border: none;
  cursor: pointer;
  font-family: inherit;
}

.nav-dropdown-wrapper:hover > .nav-dropdown-menu {
  display: block;
}

.nav-dropdown-menu {
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  background: #ffffff;
  min-width: 230px;
  z-index: 999;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.13);
  border-top: 3px solid #5F9E42;
}

/* ── Sub-dropdown ── */
.subdropdown-wrapper {
  position: relative;
}

.subdropdown-trigger {
  background: none;
  border: none;
  cursor: pointer;
  font-family: inherit;
}

.subdropdown-wrapper:hover > .subdropdown-menu {
  display: block;
}

.subdropdown-menu {
  display: none;
  position: absolute;
  left: 100%;
  top: 0;
  background: #ffffff;
  min-width: 210px;
  z-index: 999;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.13);
  border-top: 3px solid #5F9E42;
}

/* ── Item dropdown ── */
.dropdown-item {
  display: block;
  width: 100%;
  padding: 11px 18px;
  font-size: 11px;
  color: #555555;
  border-bottom: 1px solid #f0f0f0;
  text-decoration: none;
  text-align: left;
  transition: background 0.15s, color 0.15s, padding-left 0.15s;
}

.dropdown-item:hover {
  background: #f0f8ea;
  color: #5F9E42;
  padding-left: 22px;
}

.dropdown-item--active {
  background: #f0f8ea;
  color: #5F9E42;
  font-weight: 700;
}

/* ── Mobile overlay ── */
.mobile-overlay {
  display: none;
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  z-index: 998;
}

/* ── Mobile sidebar ── */
.mobile-sidebar {
  display: none;
  position: fixed;
  top: 0;
  right: 0;
  width: 300px;
  max-width: 85vw;
  height: 100vh;
  background: #fff;
  z-index: 1000;
  flex-direction: column;
  box-shadow: -4px 0 24px rgba(0, 0, 0, 0.15);
}

.mobile-sidebar-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-bottom: 1px solid #eee;
  flex-shrink: 0;
}

.mobile-close-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #333;
  padding: 4px;
}

.mobile-nav-scroll {
  flex: 1;
  overflow-y: auto;
  padding: 8px 0;
}

.mobile-link {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  padding: 13px 20px;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 1.2px;
  color: #333;
  text-decoration: none;
  text-transform: uppercase;
  border-bottom: 1px solid #f5f5f5;
  background: none;
  border-left: 3px solid transparent;
  cursor: pointer;
  text-align: left;
  transition: background 0.15s, color 0.15s;
}

.mobile-link:hover,
.mobile-link:active {
  background: #f0f8ea;
  color: #5F9E42;
}

.mobile-link--active {
  color: #5F9E42;
  background: #f0f8ea;
  border-left-color: #5F9E42;
}

.mobile-link--toggle {
  font-size: 12px;
}

.mobile-toggle-icon {
  font-size: 11px;
  transition: transform 0.25s;
  color: #999;
}

.mobile-toggle-icon.is-open {
  transform: rotate(90deg);
}

.mobile-submenu {
  padding: 0;
}

.mobile-submenu--inner {
  padding-left: 16px;
}

.mobile-sublink {
  display: block;
  padding: 11px 20px 11px 36px;
  font-size: 11.5px;
  color: #555;
  text-decoration: none;
  border-bottom: 1px solid #f5f5f5;
  transition: background 0.15s, color 0.15s;
}

.mobile-sublink:hover,
.mobile-sublink:active {
  background: #f0f8ea;
  color: #5F9E42;
}

/* ── Transitions ── */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
.slide-enter-from,
.slide-leave-to {
  transform: translateX(100%);
}

.expand-enter-active,
.expand-leave-active {
  transition: max-height 0.3s ease, opacity 0.25s ease;
  overflow: hidden;
}
.expand-enter-from,
.expand-leave-to {
  max-height: 0;
  opacity: 0;
}
.expand-enter-to,
.expand-leave-from {
  max-height: 600px;
  opacity: 1;
}

/* ── Responsive ── */
@media (max-width: 1024px) {
  .desktop-menu {
    display: none;
  }

  .hamburger-btn {
    display: flex;
  }

  .mobile-overlay {
    display: block;
  }

  .mobile-sidebar {
    display: flex;
  }
}
</style>
