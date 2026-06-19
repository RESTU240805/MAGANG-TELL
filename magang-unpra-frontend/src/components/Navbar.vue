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

        <RouterLink
          to="/"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/') }"
        >
          HOME
        </RouterLink>

        <RouterLink
          to="/our-company"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/our-company') }"
        >
          OUR COMPANY
        </RouterLink>

        <RouterLink
          to="/product"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/product') }"
        >
          PRODUCT
        </RouterLink>

        <RouterLink
          to="/news"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/news') }"
        >
          NEWS
        </RouterLink>

        <!-- Sustainability (dropdown) -->
        <div class="relative group">
          <button
            class="nav-link w-full border-none bg-transparent cursor-pointer"
            :class="{ 'nav-link--active': isActive('/sustainability') }"
          >
            SUSTAINABILITY
            <span style="font-size:9px; margin-top:1px;">▾</span>
          </button>

          <div class="dropdown-menu">
            <RouterLink
              to="/sustainability/forest-management"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/forest-management') }"
            >
              Forest Management
            </RouterLink>
            <RouterLink
              to="/sustainability/people-development"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/people-development') }"
            >
              People Development
            </RouterLink>
            <RouterLink
              to="/sustainability/supply-chain"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/supply-chain') }"
            >
              Supply Chain
            </RouterLink>
            <RouterLink
              to="/sustainability/pulp-process"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/pulp-process') }"
            >
              Pulp Process
            </RouterLink>
            <RouterLink
              to="/sustainability/safety-health"
              class="dropdown-item"
              :class="{ 'dropdown-item--active': isActive('/sustainability/safety-health') }"
            >
              Safety &amp; Health
            </RouterLink>

            <div class="relative group/csr">
              <button
                class="dropdown-item w-full text-left border-none bg-transparent cursor-pointer flex justify-between items-center"
                :class="{ 'dropdown-item--active': isActive('/sustainability/csr') }"
              >
                Corporate Social Responsibility
                <span style="color:#aaa; margin-left:8px;">›</span>
              </button>
              <div class="subdropdown-menu">
                <RouterLink
                  to="/sustainability/csr/vision"
                  class="dropdown-item"
                  :class="{ 'dropdown-item--active': isActive('/sustainability/csr/vision') }"
                >
                  Vision And Mission
                </RouterLink>
                <RouterLink
                  to="/sustainability/csr/community"
                  class="dropdown-item"
                  :class="{ 'dropdown-item--active': isActive('/sustainability/csr/community') }"
                >
                  Local Community Development
                </RouterLink>
                <RouterLink
                  to="/sustainability/csr/report"
                  class="dropdown-item"
                  style="border-bottom: none;"
                  :class="{ 'dropdown-item--active': isActive('/sustainability/csr/report') }"
                >
                  CSR Report
                </RouterLink>
              </div>
            </div>
          </div>
        </div>

        <!-- Biodiversity -->
        <a
          href="https://kehati.telpp.com/"
          target="_blank"
          rel="noopener noreferrer"
          class="nav-link"
        >
          BIODIVERSITY
        </a>

        <!-- E-Proc -->
        <a
          href="https://eproc.telpp.com/_gst_home.php"
          target="_blank"
          rel="noopener noreferrer"
          class="nav-link"
        >
          E-PROC
        </a>

        <!-- E-Recruitment -->
        <a
          href="https://erecruitment.telpp.com/er_tlpp/"
          target="_blank"
          rel="noopener noreferrer"
          class="nav-link"
        >
          E-RECRUITMENT
        </a>

        <!-- Contact -->
        <RouterLink
          to="/contact"
          class="nav-link"
          :class="{ 'nav-link--active': isActive('/contact') }"
        >
          CONTACT
        </RouterLink>
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
          <RouterLink to="/" class="mobile-link" :class="{ 'mobile-link--active': isActive('/') }" @click="mobileOpen = false">
            HOME
          </RouterLink>
          <RouterLink to="/our-company" class="mobile-link" :class="{ 'mobile-link--active': isActive('/our-company') }" @click="mobileOpen = false">
            OUR COMPANY
          </RouterLink>
          <RouterLink to="/product" class="mobile-link" :class="{ 'mobile-link--active': isActive('/product') }" @click="mobileOpen = false">
            PRODUCT
          </RouterLink>
          <RouterLink to="/news" class="mobile-link" :class="{ 'mobile-link--active': isActive('/news') }" @click="mobileOpen = false">
            NEWS
          </RouterLink>

          <!-- Sustainability toggle -->
          <button
            class="mobile-link mobile-link--toggle"
            :class="{ 'mobile-link--active': isActive('/sustainability') }"
            @click="sustainabilityOpen = !sustainabilityOpen"
          >
            SUSTAINABILITY
            <span class="mobile-toggle-icon" :class="{ 'is-open': sustainabilityOpen }">▾</span>
          </button>
          <Transition name="expand">
            <div v-if="sustainabilityOpen" class="mobile-submenu">
              <RouterLink to="/sustainability/forest-management" class="mobile-sublink" @click="mobileOpen = false">Forest Management</RouterLink>
              <RouterLink to="/sustainability/people-development" class="mobile-sublink" @click="mobileOpen = false">People Development</RouterLink>
              <RouterLink to="/sustainability/supply-chain" class="mobile-sublink" @click="mobileOpen = false">Supply Chain</RouterLink>
              <RouterLink to="/sustainability/pulp-process" class="mobile-sublink" @click="mobileOpen = false">Pulp Process</RouterLink>
              <RouterLink to="/sustainability/safety-health" class="mobile-sublink" @click="mobileOpen = false">Safety & Health</RouterLink>

              <!-- CSR toggle -->
              <button
                class="mobile-sublink mobile-link--toggle"
                @click="csrOpen = !csrOpen"
              >
                Corporate Social Responsibility
                <span class="mobile-toggle-icon" :class="{ 'is-open': csrOpen }">›</span>
              </button>
              <Transition name="expand">
                <div v-if="csrOpen" class="mobile-submenu mobile-submenu--inner">
                  <RouterLink to="/sustainability/csr/vision" class="mobile-sublink" @click="mobileOpen = false">Vision And Mission</RouterLink>
                  <RouterLink to="/sustainability/csr/community" class="mobile-sublink" @click="mobileOpen = false">Local Community Development</RouterLink>
                  <RouterLink to="/sustainability/csr/report" class="mobile-sublink" @click="mobileOpen = false">CSR Report</RouterLink>
                </div>
              </Transition>
            </div>
          </Transition>

          <a href="https://kehati.telpp.com/" target="_blank" rel="noopener noreferrer" class="mobile-link" @click="mobileOpen = false">
            BIODIVERSITY
          </a>
          <a href="https://eproc.telpp.com/_gst_home.php" target="_blank" rel="noopener noreferrer" class="mobile-link" @click="mobileOpen = false">
            E-PROC
          </a>
          <a href="https://erecruitment.telpp.com/er_tlpp/" target="_blank" rel="noopener noreferrer" class="mobile-link" @click="mobileOpen = false">
            E-RECRUITMENT
          </a>
          <RouterLink to="/contact" class="mobile-link" :class="{ 'mobile-link--active': isActive('/contact') }" @click="mobileOpen = false">
            CONTACT
          </RouterLink>
        </div>
      </div>
    </Transition>
  </nav>
</template>

<script setup>
import { ref, watch } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()
const mobileOpen = ref(false)
const sustainabilityOpen = ref(false)
const csrOpen = ref(false)

const isActive = (path) => {
  if (path === '/') {return route.path === '/'}
  return route.path.startsWith(path)
}

const toggleMobile = () => {
  mobileOpen.value = !mobileOpen.value
}

watch(() => route.path, () => {
  mobileOpen.value = false
  sustainabilityOpen.value = false
  csrOpen.value = false
})
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

/* ── Dropdown utama ── */
.dropdown-menu {
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

.group:hover > .dropdown-menu {
  display: block;
}

/* ── Sub-dropdown CSR ── */
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

.group\/csr:hover > .subdropdown-menu {
  display: block;
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
