[
    {
        "shape_id": "1311-0", // Identificador da forma (rota) da viagem
        "stop_sequence": 20, // Sequência numérica da parada
        "trip_headsign": "Ulysses via José Loureiro - Ida para Estação Guanabara", // Destino da viagem
        "stop_time": "14:20:37", // Horário previsto de chegada à parada
        "time_left": 2259, // Tempo restante em segundos para a chegada
        "distance_diff": null, // Diferença de distância (não está preenchido nesse exemplo)
        "next_trip": 1, // Próxima viagem relacionada
        "trips": [ // Lista de informações das viagens
            {
                "eta": "13:31:41", // Horário estimado de chegada
                "time_left": -676, // Tempo restante em segundos para a chegada (negativo indica que o ônibus já passou)
                "start_time_diff": -1377, // Diferença de tempo desde o início da viagem
                "trip_id": "G-DOM-10-25-1311-0", // Identificador da viagem
                "trip_status": "LIVE", // Status da viagem
                "report_lat": -26.319776, // Latitude do relatório do ônibus
                "report_lon": -48.821276, // Longitude do relatório do ônibus
                "stop_name": "Estação Guanabara - Plataforma 1", // Nome da parada
                "stop_distance": -176.65, // Distância em metros entre o ônibus e a parada (negativo indica aproximação)
                "stop_order": 35, // Ordem da parada na viagem
                "vehicle_prefix": "11502", // Prefixo do veículo
                "distance_diff": -5004, // Diferença de distância (em metros)
                "report_time_diff": 14 // Diferença de tempo desde o último relatório
            },
            // ... Outras viagens com campos similares ...
        ]
    }
]