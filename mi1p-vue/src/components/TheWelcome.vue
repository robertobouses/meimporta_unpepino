<script setup>
import WelcomeItem from './WelcomeItem.vue'
import BusquedasIcon from './icons/IconBusquedas.vue'
import CalculadorIcon from './icons/IconCalculador.vue'
import DefineTuCultivoIcon from './icons/IconDefineTuCultivo.vue'
import PlagasIcon from './icons/IconPlagas.vue'
import LecturasIcon from './icons/IconLecturas.vue'
import 'bootstrap-icons/font/bootstrap-icons.css';
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <BusquedasIcon />
    </template>
    <template #heading>Búsquedas</template>

    Encuentra la información sobre tu cultivo, o filtra y encuentra el cultivo que ni sospechabas necesitar, nuestros filtros y algoritmos trabajan para ti!
    
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CalculadorIcon />
    </template>
    <template #heading>Calculador de plantas</template>

    Consigue saber las plantas que necesitas para tu cultivo, o los metros que necesitas para tus plantines, obtén la distancia de plantación...
   
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <DefineTuCultivoIcon />
    </template>
    <template #heading>Define tu propio cultivo</template>

    Un espacio muy tuyo, introduce los datos de tu plantación y tendrás en la palma de tu mano el diario de cultivo que necesitas. No te olvidarás de regar más, sabrás cuándo plantar, cuánto fertilizar exactamente y tendrás hasta predicciones de cuantos frutos vas a obtener. Todo adaptado a tu cultivo real.
    
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <PlagasIcon />
    </template>
    <template #heading>Plagas y problemas</template>

    Porque nada es perfecto, los imprevistos y la naturaleza son hermanos, la mejor arma la sabiduría, conoce todo tipo de plagas que pueden afectar a tu cultivo, aprende también sobre otros problemas meteorológicos, hongos, de nutrición e incluso físicos. Cuidamos de tus plantas.
    
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <LecturasIcon />
    </template>
    <template #heading>Lecturas</template>

    Si eres curioso, aún tenemos más, lee curiosidades, sobre modos de cultivo, un poco de historia y sobre todo sobre agricultura.
   </WelcomeItem>
</template>
