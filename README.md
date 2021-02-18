# mutants
Mutants challenge

La API se encuentra publicada en 3.89.20.240 en el puerto 8000

Enpoints: 

 - POST http://18.206.244.132:8000/mutant/
 
    Devuelve si un adn pertenece a un mutante o no

    El json debe de tener el formato:
    
        {
            "dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
        }
    
    En el caso de verificar un mutante retorna HTTP 200-OK y en el caso contrario 403-Forbidden

- GET http://18.206.244.132:8000/stats

    Retorna las estadísitcas de las verificaciones de ADN
    
    El formato de respuesta es el siguiente: 
    
        {"count_mutant_dna":40, "count_human_dna":100: "ratio":0.4}