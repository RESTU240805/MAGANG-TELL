<template>
  <section class="relative h-[420px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('/images/lokasi pabrik.jpeg')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
    <div class="absolute top-10 left-10 flex items-center gap-2 opacity-0 anim-fade-right" style="animation-delay: 0.1s">
      <span class="w-1 h-5 bg-green-400 rounded-full inline-block"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>
    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4 opacity-0 anim-fade-up" style="animation-delay: 0.2s">
        Pulp <span class="text-green-400">Process</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed opacity-0 anim-fade-up" style="animation-delay: 0.3s">
        Explore the complete industrial pulp production process from raw logs to finished bales.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400 opacity-0 anim-fade-up" style="animation-delay: 0.4s">
        <RouterLink to="/" class="hover:text-white cursor-pointer transition">Home</RouterLink>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">Pulp Process</span>
      </div>
    </div>
  </section>

  <!-- 3D Interactive Map -->
  <section class="bg-[#0a1208] py-10">
    <div class="max-w-7xl mx-auto px-4">
      <div class="anim-item rounded-xl overflow-hidden shadow-2xl border border-white/10">
        <PulpProcess3D />
      </div>
    </div>
  </section>

  <!-- Process Description -->
  <section class="py-16 bg-white">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-3xl font-black text-green-600 mb-6 anim-item">Pulp Process</h2>
      <p class="text-gray-600 text-sm leading-relaxed mb-6 anim-item">
        There are two main processes in the Pulp and Paper Mill, there are:
      </p>

      <div class="space-y-6">
        <div class="bg-green-50 rounded-xl p-6 border border-green-100 anim-item">
          <h3 class="text-lg font-bold text-gray-900 mb-2">1. Pulp Making Process</h3>
          <p class="text-gray-600 text-sm leading-relaxed">
            Pulp making process is represents the process by which wood or other lignocellulosic material
            that reduced to be fibrous mass. Pulp is produced at PT. TELPP is 450,000 Adt (BKP) / year
            or 1,430 Adt/BKPd.
          </p>
        </div>

        <div class="bg-red-50 rounded-xl p-6 border border-red-100 anim-item">
          <h3 class="text-lg font-bold text-gray-900 mb-2">2. Chemical Recovery Process</h3>
          <p class="text-gray-600 text-sm leading-relaxed">
            Chemical recovery process is the process to recover again cooking chemical (White Liquor)
            that used to cooked wood chip become pulp. White Liquor contains active chemical NaOH and
            Na&#8322;S. Design capacity for Chemical recovery process is 2400 TDS/d solid burning in Recovery Boiler.
          </p>
        </div>
      </div>
    </div>
  </section>

  <!-- Pulp Process Flow Diagram -->
  <section class="py-12 bg-gray-50">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-2xl font-black text-gray-900 text-center mb-8 anim-item">
        Pulp Process <span class="text-green-600">Flow</span>
      </h2>
      <div class="anim-item bg-white rounded-2xl overflow-hidden shadow-sm border border-gray-100">
        <img src="/pulp proces/all_process-2.jpg" alt="Pulp Process Flow Diagram"
          class="w-full h-auto" />
      </div>
    </div>
  </section>

  <!-- Pulp Process Accordion -->
  <section class="py-16 bg-white">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-3xl font-black text-green-600 mb-8 anim-item">Pulp Process</h2>

      <div class="space-y-3">
        <div v-for="(section, i) in pulpSections" :key="i"
          class="border border-gray-200 rounded-xl overflow-hidden anim-item"
          :style="{ animationDelay: `${i * 0.08}s` }">
          <button @click="togglePulpSection(i)"
            class="w-full flex items-center justify-between px-6 py-4 text-left transition-colors cursor-pointer"
            :class="section.open ? 'bg-green-600 text-white' : 'bg-green-600 text-white hover:bg-green-700'">
            <span class="font-bold text-sm">{{ section.title }}</span>
            <svg class="w-5 h-5 transition-transform duration-300" :class="section.open ? 'rotate-180' : ''"
              fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>
          <div v-show="section.open" class="px-6 py-5 bg-gray-50">
            <img v-if="section.image" :src="section.image" :alt="section.title"
              class="w-full h-72 object-cover rounded-lg mb-4" />
            <p class="text-gray-600 text-sm leading-relaxed">{{ section.desc }}</p>
            <div v-if="section.details" class="mt-4 space-y-2">
              <div v-for="(detail, j) in section.details" :key="j"
                class="flex items-start gap-3 text-sm">
                <span class="w-1.5 h-1.5 rounded-full bg-green-500 mt-1.5 shrink-0"></span>
                <span class="text-gray-600">{{ detail }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- Warehouse Photo -->
  <section class="py-12 bg-gray-50">
    <div class="anim-item overflow-hidden shadow-lg">
      <img src="/pulp proces/IMG_6725.jpg" alt="Finished Pulp Bales Warehouse"
        class="w-full h-[500px] object-cover" />
    </div>
  </section>

  <!-- Chemical Recovery Process -->
  <section class="py-16 bg-white">
    <div class="max-w-4xl mx-auto px-10">
      <h2 class="text-3xl font-black text-green-600 mb-4 anim-item">Chemical Recovery Process</h2>
      <p class="text-gray-600 text-sm leading-relaxed mb-8 anim-item">
        The chemical recovery cycle regenerates cooking chemicals and generates energy.
        Black liquor from the washing process is concentrated and burned to recover
        inorganic chemicals while producing steam for power generation.
      </p>

      <div class="flex flex-wrap gap-2 mb-8 anim-item">
        <button v-for="(tab, i) in recoveryTabs" :key="i"
          @click="activeRecoveryTab = i"
          class="px-4 py-2 rounded-full text-xs font-bold tracking-wide transition-all cursor-pointer border"
          :class="activeRecoveryTab === i
            ? 'bg-green-600 text-white border-green-600'
            : 'bg-white text-gray-600 border-gray-200 hover:border-green-400 hover:text-green-600'">
          {{ tab }}
        </button>
      </div>

      <div class="bg-gray-50 rounded-xl overflow-hidden border border-gray-100 anim-item">
        <img v-if="recoveryData[activeRecoveryTab].image"
          :src="recoveryData[activeRecoveryTab].image"
          :alt="recoveryData[activeRecoveryTab].title"
          class="w-full h-72 object-cover" />
        <div class="p-6">
          <h3 class="text-lg font-bold text-gray-900 mb-3">{{ recoveryData[activeRecoveryTab].title }}</h3>
          <p class="text-gray-600 text-sm leading-relaxed mb-4">{{ recoveryData[activeRecoveryTab].desc }}</p>
          <div class="grid grid-cols-2 gap-4">
            <div v-for="(spec, j) in recoveryData[activeRecoveryTab].specs" :key="j"
              class="bg-white rounded-lg p-4 border border-gray-100">
              <div class="text-xs text-gray-400 mb-1">{{ spec.label }}</div>
              <div class="text-sm font-bold text-gray-800">{{ spec.value }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <footer class="site-footer">
      <div class="footer-container">
        <div class="footer-left-content"></div>
        <div class="footer-copyright">
          <p>Copyright 2026 PT TELPP. All right reserved.</p>
        </div>
      </div>
    </footer>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, nextTick } from 'vue'
import { RouterLink } from 'vue-router'
import PulpProcess3D from '@/components/PulpProcess3D.vue'

const pulpSections = reactive([
  {
    title: 'Wood and Chip Handling',
    open: true,
    image: '/pulp proces/Chipping-Proses.jpg',
    desc: 'Raw logs from sustainable industrial plantations (Acacia Mangium & Eucalyptus Pellita) are transported to the mill, stored in wet ponds, then processed through debarking and chipping.',
    details: [
      'Log Storage Ponds — logs kept wet to preserve fiber quality',
      'Debarking Drum — rotary drum strips bark (bark used as biomass fuel)',
      'Disk Chipper — reduces debarked logs into uniform 25mm chips',
      'Chip Screening — separates oversized chips (rechipped) and fines (fuel)'
    ]
  },
  {
    title: 'Fiber Line',
    open: false,
    image: '/pulp proces/FiberLine.jpg',
    desc: 'Wood chips are cooked under high pressure with white liquor to separate cellulose fibers from lignin, then washed and delignified.',
    details: [
      'Chip Bin & Digester — Kamyr continuous digester at 170\u00b0C with NaOH + Na\u2082S',
      'Blow Tank — receives cooked pulp at atmospheric pressure',
      'Brownstock Washing — multi-stage washers remove black liquor',
      'O\u2082 Delignification — reduces kappa number 40-50% using oxygen'
    ]
  },
  {
    title: 'Pulp Machine',
    open: false,
    image: '/pulp proces/pulp-machine.jpg',
    desc: 'The bleached pulp slurry is formed into sheets on a Fourdrinier wire, pressed, and steam-dried into market pulp.',
    details: [
      'Forming Section — diluted pulp (0.5-1% consistency) deposited on wire',
      'Press Section — water removed by pressing between felt rolls',
      'Dryer Section — steam-heated cylinders dry pulp to ~90% solids',
      'Production capacity: 450,000 ADt/year (1,430 Adt/day)'
    ]
  },
  {
    title: 'Warehouse',
    open: false,
    image: '/pulp proces/IMG_6725.jpg',
    desc: 'Finished pulp sheets are cut, stacked, and compressed into 250 kg bales wrapped in kraft liner for storage and export.',
    details: [
      'Bale Press — sheets compressed into 250 kg bales',
      'Wrapping — bales wrapped in kraft liner for protection',
      'Climate-controlled storage — temperature and humidity regulated',
      'Conveyor belt to jetty — bales loaded onto cargo ships for export'
    ]
  }
])

function togglePulpSection(index) {
  pulpSections[index].open = !pulpSections[index].open
}

const activeRecoveryTab = ref(0)
const recoveryTabs = ['Evaporator', 'Recovery Boiler', 'Recausticizing', 'Lime Kiln']

const recoveryData = [
  {
    title: 'Evaporator Plant (7-Effect)',
    image: '/pulp proces/Epavorator.jpg',
    desc: 'Multiple-effect evaporator concentrates weak black liquor from 15% to 72% dissolved solids for firing in the Recovery Boiler. The 7-effect system maximizes steam economy.',
    specs: [
      { label: 'Inlet Concentration', value: '15% Dissolved Solids' },
      { label: 'Outlet Concentration', value: '72% Dissolved Solids' },
      { label: 'Effect Stages', value: '7-Effect System' },
      { label: 'Heat Source', value: 'Exhaust Steam from Boilers' }
    ]
  },
  {
    title: 'Recovery Boiler',
    image: '/pulp proces/Recovery-Boiler.jpg',
    desc: 'Burns concentrated black liquor (70% DS) to recover inorganic smelt and generate high-pressure steam. The smelt dissolves in green liquor for recausticizing.',
    specs: [
      { label: 'Design Capacity', value: '2,400 TDS/day' },
      { label: 'Fuel', value: 'Concentrated Black Liquor' },
      { label: 'Products', value: 'Smelt + High-Pressure Steam' },
      { label: 'Steam Usage', value: 'Power Generation + Process' }
    ]
  },
  {
    title: 'Recausticizing Plant',
    image: '/pulp proces/Recautisizing.jpg',
    desc: 'Green liquor (Na\u2082CO\u2083 + Na\u2082S) is reacted with Ca(OH)\u2082 to regenerate White Liquor (NaOH + Na\u2082S) for the cooking cycle. Lime mud (CaCO\u2083) is sent to the lime kiln.',
    specs: [
      { label: 'Input', value: 'Green Liquor (Na\u2082CO\u2083 + Na\u2082S)' },
      { label: 'Reagent', value: 'Ca(OH)\u2082 (Slaked Lime)' },
      { label: 'Output', value: 'White Liquor (NaOH + Na\u2082S)' },
      { label: 'By-product', value: 'Lime Mud (CaCO\u2083) \u2192 Kiln' }
    ]
  },
  {
    title: 'Lime Kiln',
    image: '/pulp proces/Lime-Kiln.jpg',
    desc: 'Rotary kiln (~150m) reburns CaCO\u2083 lime mud at 1000\u00b0C to regenerate CaO quicklime for the recausticizing cycle. Completes the chemical recovery loop.',
    specs: [
      { label: 'Length', value: '~150 meters' },
      { label: 'Temperature', value: '1,000\u00b0C' },
      { label: 'Input', value: 'Lime Mud (CaCO\u2083)' },
      { label: 'Output', value: 'Quicklime (CaO) \u2192 Recausticizing' }
    ]
  }
]

let observer = null

onMounted(async () => {
  await nextTick()
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-visible')
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.15 })
  document.querySelectorAll('.anim-item').forEach(el => observer.observe(el))
})

onUnmounted(() => {
  if (observer) {observer.disconnect()}
})
</script>

<style scoped>
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
.anim-item {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-item.anim-visible {
  opacity: 1;
  transform: translateY(0);
}

/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  /* Warna hijau tua disesuaikan dengan identitas PT TEL pada gambar */
  background-color: #5F9E42; 
  padding: 25px 0; /* Memberikan ruang tinggi baris yang pas */
  margin-top: 40px; /* Jarak pemisah dari konten di atasnya */
}

.footer-container {
  display: flex;
  align-items: center;
  justify-content: center; /* DIUBAH: Membuat konten berkumpul di tengah */
  max-width: 1140px; /* Menyelaraskan lebar dengan konten atas website */
  margin: 0 auto;
  padding: 0 40px; /* Jarak aman agar teks tidak terlalu menempel ke pinggir layar */
}

/* Kolom Kiri Footer */
.footer-left-content {
  display: none; /* DIUBAH: Disembunyikan karena kosong, agar tidak memakan space */
}

/* Kolom Kanan Footer (Teks Copyright) */
.footer-copyright {
  text-align: center; /* DIUBAH: Memastikan teks rata tengah */
  width: 100%; /* Memastikan block mengambil ruang penuh untuk fungsi centering */
}

.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px; /* Ukuran teks kecil, tipis, dan normal untuk footer */
  color: #ffffff;  /* Warna teks putih bersih agar kontras dengan background hijau */
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
</style>
