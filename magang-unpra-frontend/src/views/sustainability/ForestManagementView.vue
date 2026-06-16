<template>

  <!-- ── HERO ── -->
  <section class="relative h-[420px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('/images/hutan.jpg')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>

    <!-- Label atas kiri -->
    <div class="absolute top-10 left-10 flex items-center gap-2 opacity-0 anim-fade-right" style="animation-delay: 0.1s">
      <span class="w-1 h-5 bg-green-400 rounded-full inline-block"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <!-- Teks bawah -->
    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4 opacity-0 anim-fade-up" style="animation-delay: 0.2s">
        Forest <span class="text-green-400">Management</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed opacity-0 anim-fade-up" style="animation-delay: 0.3s">
        Sustainable and responsible forest management for a greener future in South Sumatra.
      </p>
      <!-- Breadcrumb -->
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400 opacity-0 anim-fade-up" style="animation-delay: 0.4s">
        <RouterLink to="/" class="hover:text-white cursor-pointer transition">Home</RouterLink>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">Forest Management</span>
      </div>
    </div>
  </section>

  <!-- ── INTRO ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <!-- Teks -->
      <div class="anim-slide-right mb-10">
        <span class="inline-block bg-green-50 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-4">
          Our Approach
        </span>
        <h2 class="text-3xl font-black text-gray-900 leading-tight mb-5">
          High Quality Pulp from <span class="text-green-600">100% Sustainable Wood</span>
        </h2>
        <div class="space-y-4 text-gray-500 text-sm leading-relaxed">
          <p>
             PT Tanjungenim Lestari Pulp and Paper (TEL) produces high quality pulp from 100% wood
            based on industrial plantations sourced from a group company, PT Musi Hutan Persada (MHP),
            a highly dedicated Industrial Plantation Company with 296,400 ha of concession area in South Sumatra.
             MHP runs its operations with Sustainable and Sustainable Forest Management.
          </p>
          <p>
           Pulp is produced from two types of wood, namely Acacia Mangium (AM) and Eucalyptus Pellita (EP). Since the start of the plant's operation in 1999, the main product is the Acacia Mangium Pulp and then through continuous efforts, TEL is working with MHP to develop new alternative raw materials, until now Eucalyptus Pellita Pulp has become the main product.
          </p>
        </div>

        <div class="mt-6 border-l-4 border-green-500 pl-4 bg-green-50 py-3 pr-4 rounded-r-xl">
          <p class="text-green-800 text-sm font-medium leading-relaxed">
            Timber delivery from the Industrial Plantation area to the Factory Site continues 365 days per year by truck through a special road network built by MHP. TEL sets standards for the size and quality of logs that must be met by MHP.
          </p>
        </div>
      </div>

      <!-- Gambar Slider Full Width -->
      <div class="relative anim-slide-left">
        <div class="relative rounded-2xl overflow-hidden shadow-xl h-[480px] group">
          <img
            v-for="(slide, i) in heroSlides"
            :key="i"
            :src="slide"
            alt="Forest Management"
            class="absolute inset-0 w-full h-full object-cover"
            :style="{ opacity: currentSlide === i ? 1 : 0, transition: 'opacity 1s ease-in-out' }"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>

          <!-- Arrows -->
          <button @click="manualHeroNav(prevHeroSlide)"
            class="absolute left-3 top-1/2 -translate-y-1/2 w-9 h-9 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer hover:bg-white/40">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
          </button>
          <button @click="manualHeroNav(nextHeroSlide)"
            class="absolute right-3 top-1/2 -translate-y-1/2 w-9 h-9 rounded-full bg-white/20 backdrop-blur-md text-white border border-white/30 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer hover:bg-white/40">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          </button>

          <!-- Dots -->
          <div class="absolute bottom-3 left-1/2 -translate-x-1/2 flex gap-2 z-10">
            <button v-for="(_, i) in heroSlides" :key="i" @click="manualHeroNav(() => goToHeroSlide(i))"
              class="h-1.5 rounded-full transition-all duration-500 cursor-pointer"
              :class="currentSlide === i ? 'w-8 bg-white' : 'w-3 bg-white/40 hover:bg-white/60'">
            </button>
          </div>
        </div>
      </div>

    </div>
  </section>

  <!-- ── WOOD TYPES ── -->
  <section class="py-20 bg-[#14532d]">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-12">
        <span class="inline-block bg-green-800 text-green-200 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3 anim-item">
          Raw Materials
        </span>
        <h2 class="text-3xl font-black text-white anim-item">Types of Wood Used</h2>
        <p class="text-green-200/60 text-sm mt-2 max-w-md mx-auto anim-item">
          Pulp is produced from two types of wood sourced from sustainable industrial plantations.
        </p>
      </div>

      <div class="grid grid-cols-2 gap-8">

        <!-- Acacia -->
        <div class="bg-white/10 backdrop-blur-sm rounded-2xl overflow-hidden border border-white/10 hover:bg-white/15 transition group anim-item" style="animation-delay: 0.1s">
          <div class="h-48 overflow-hidden">
            <img
              src="/images/accsia.jpg"
              alt="Acacia"
              class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-700"
            />
          </div>
          <div class="p-6">
            <div class="flex items-center gap-2 mb-2">
              <span class="w-2 h-2 rounded-full bg-green-400 inline-block"></span>
              <span class="text-xs font-bold text-green-300 tracking-widest uppercase">Initial Product</span>
            </div>
            <h3 class="text-lg font-black text-white mb-2">Acacia Mangium (AM)</h3>
            <p class="text-green-200/60 text-sm leading-relaxed">
              The main product since the start of TEL's plant operation in 1999.
              Acacia Mangium Pulp served as the foundation of production.
            </p>
          </div>
        </div>

        <!-- Eucalyptus -->
        <div class="bg-white/10 backdrop-blur-sm rounded-2xl overflow-hidden border border-white/10 hover:bg-white/15 transition group anim-item" style="animation-delay: 0.2s">
          <div class="h-48 overflow-hidden">
            <img
              src="/images/eccalyptus.jpg"
              alt="Eucalyptus"
              class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-700"
            />
          </div>
          <div class="p-6">
            <div class="flex items-center gap-2 mb-2">
              <span class="w-2 h-2 rounded-full bg-emerald-400 inline-block"></span>
              <span class="text-xs font-bold text-emerald-300 tracking-widest uppercase">Main Product Now</span>
            </div>
            <h3 class="text-lg font-black text-white mb-2">Eucalyptus Pellita (EP)</h3>
            <p class="text-green-200/60 text-sm leading-relaxed">
              Through continuous innovation and development, Eucalyptus Pellita Pulp
              has now become the primary raw material used by TEL.
            </p>
          </div>
        </div>

      </div>
    </div>
  </section>

  

</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { RouterLink } from 'vue-router'

const deliveryFeatures = [
  'Special road network built and maintained by MHP',
  'Strict log size and quality standards set by TEL',
  'Year-round continuous supply to factory site',
  'Efficient truck-based transportation system',
]

// ─── Slider ───────────────────────────────────────────────────
const currentSlide = ref(0)
let heroTimer = null

const heroSlides = [
  '/images/hutan.jpg',
  '/images/hutan-960x400.jpg',
  '/images/MHP4-960x400.jpg',
  '/images/MHP2-940x392.jpg',
]

const nextHeroSlide = () => { currentSlide.value = (currentSlide.value + 1) % heroSlides.length }
const prevHeroSlide = () => { currentSlide.value = (currentSlide.value - 1 + heroSlides.length) % heroSlides.length }
const goToHeroSlide = (i) => { currentSlide.value = i }
const startHeroTimer = () => { heroTimer = setInterval(nextHeroSlide, 5000) }
const stopHeroTimer = () => { clearInterval(heroTimer) }
const manualHeroNav = (fn) => { stopHeroTimer(); fn(); startHeroTimer() }

// ─── Scroll Observer ──────────────────────────────────────────
let observer = null

onMounted(async () => {
  await nextTick()
  initObserver()
  startHeroTimer()
})

onUnmounted(() => {
  if (observer) observer.disconnect()
  stopHeroTimer()
})

const initObserver = () => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-visible')
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.15 })

  document.querySelectorAll('.anim-item, .anim-slide-right, .anim-slide-left').forEach(el => {
    observer.observe(el)
  })
}
</script>

<style scoped>
/* Slider */
.slider-img {
  transition: opacity 1s ease-in-out;
}

/* Hero animations */
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}
.anim-fade-up {
  animation: fadeUp 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

@keyframes fadeRight {
  from { opacity: 0; transform: translateX(-24px); }
  to { opacity: 1; transform: translateX(0); }
}
.anim-fade-right {
  animation: fadeRight 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

/* Scroll animations */
.anim-item {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-item.anim-visible {
  opacity: 1;
  transform: translateY(0);
}

.anim-slide-right {
  opacity: 0;
  transform: translateX(-40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-right.anim-visible {
  opacity: 1;
  transform: translateX(0);
}

.anim-slide-left {
  opacity: 0;
  transform: translateX(40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-left.anim-visible {
  opacity: 1;
  transform: translateX(0);
}
</style>
