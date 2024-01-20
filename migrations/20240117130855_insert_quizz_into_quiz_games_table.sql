-- +goose Up
-- +goose StatementBegin
INSERT INTO quiz_games (gid, data, created_at, updated_at)
values ((SELECT id FROM games WHERE name = 'Quiz' AND game_version = '1.0'),
        '{
    "topic": [
        {
            "id": 1,
            "name": "Energie Solaire",
            "description" : "L''énergie solaire exploite la puissance du soleil pour produire de l''électricité de manière propre et renouvelable.",
            "quizzes": [
                {
                    "id": 1,
                    "topic": "Energie Solaire",
                    "title": "Quiz Aléatoire sur l''Énergie Solaire",
                    "questions": [
                        {
                            "question": "Quel est le processus par lequel l''énergie solaire est convertie en électricité ?",
                            "good_answer": "La conversion photovoltaïque",
                            "answers": [
                                "La photosynthèse",
                                "La combustion solaire",
                                "La conversion photovoltaïque"
                            ]
                        },
                        {
                            "question": "Quel élément est principalement utilisé dans les panneaux solaires pour capter la lumière du soleil ?",
                            "good_answer": "Le silicium",
                            "answers": [
                                "L''hydrogène",
                                "Le carbone",
                                "Le silicium"
                            ]
                        },
                        {
                            "question": "Quel est l''appareil qui permet de suivre la trajectoire du soleil pour maximiser la capture d''énergie solaire ?",
                            "good_answer": "Le suiveur solaire",
                            "answers": [
                                "Le réflecteur solaire",
                                "L''accumulateur solaire",
                                "Le suiveur solaire"
                            ]
                        },
                        {
                            "question": "Quelle est l''unité de mesure de la puissance des panneaux solaires ?",
                            "good_answer": "Watt-crête (Wc)",
                            "answers": [
                                "Lumen (lm)",
                                "Kilowatt-heure (kWh)",
                                "Watt-crête (Wc)"
                            ]
                        },
                        {
                            "question": "Quel pays est le plus grand producteur d''énergie solaire au monde ?",
                            "good_answer": "La Chine",
                            "answers": [
                                "Les États-Unis",
                                "L''Inde",
                                "La Chine"
                            ]
                        },
                        {
                            "question": "Quel avantage majeur de l''énergie solaire en fait une source d''énergie renouvelable attrayante ?",
                            "good_answer": "Elle ne produit pas de pollution atmosphérique",
                            "answers": [
                                "Elle est bon marché",
                                "Elle est facile à stocker",
                                "Elle ne produit pas de pollution atmosphérique"
                            ]
                        },
                        {
                            "question": "Quel dispositif permet de stocker l''énergie solaire pour une utilisation ultérieure ?",
                            "good_answer": "Les batteries solaires",
                            "answers": [
                                "Les panneaux solaires",
                                "Les onduleurs solaires",
                                "Les batteries solaires"
                            ]
                        },
                        {
                            "question": "Quelle est l''efficacité typique des panneaux solaires commerciaux ?",
                            "good_answer": "Entre 15% et 20%",
                            "answers": [
                                "Plus de 50%",
                                "Moins de 5%",
                                "Entre 15% et 20%"
                            ]
                        }
                    ]
                },
                {
                    "id": 2,
                    "topic": "Energie Solaire",
                    "title": "Le Quiz Solaire Incroyable",
                    "questions": [
                        {
                            "question": "Quelle est la source d''énergie principale du solaire photovoltaïque ?",
                            "good_answer": "Lumière du soleil",
                            "answers": [
                                "Vent",
                                "Pétrole",
                                "Lumière du soleil"
                            ]
                        },
                        {
                            "question": "Quel composant clé est utilisé pour convertir la lumière en électricité dans les panneaux solaires ?",
                            "good_answer": "Cellules photovoltaïques",
                            "answers": [
                                "Batteries",
                                "Aimants",
                                "Cellules photovoltaïques"
                            ]
                        },
                        {
                            "question": "Quel pays est le plus grand producteur d''énergie solaire au monde ?",
                            "good_answer": "Chine",
                            "answers": [
                                "États-Unis",
                                "Inde",
                                "Chine"
                            ]
                        },
                        {
                            "question": "Comment est stockée l''énergie solaire pour une utilisation ultérieure ?",
                            "good_answer": "Batteries",
                            "answers": [
                                "Réservoirs d''eau",
                                "Sacs de sable",
                                "Batteries"
                            ]
                        },
                        {
                            "question": "Quel est l''angle optimal pour l''installation de panneaux solaires afin de maximiser l''exposition au soleil ?",
                            "good_answer": "Entre 30 et 45 degrés",
                            "answers": [
                                "90 degrés",
                                "0 degrés",
                                "Entre 30 et 45 degrés"
                            ]
                        },
                        {
                            "question": "Quel type de rayonnement solaire est converti en énergie électrique par les panneaux solaires ?",
                            "good_answer": "Rayonnement solaire photovoltaïque",
                            "answers": [
                                "Rayonnement infrarouge",
                                "Rayonnement ultraviolet",
                                "Rayonnement solaire photovoltaïque"
                            ]
                        },
                        {
                            "question": "Quelle est la durée de vie typique des panneaux solaires ?",
                            "good_answer": "Environ 25 ans",
                            "answers": [
                                "5 ans",
                                "50 ans",
                                "Environ 25 ans"
                            ]
                        },
                        {
                            "question": "Quelle est la principale source d''énergie utilisée pour chauffer l''eau à l''aide de l''énergie solaire ?",
                            "good_answer": "Énergie solaire thermique",
                            "answers": [
                                "Gaz naturel",
                                "Électricité",
                                "Énergie solaire thermique"
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "id": 2,
            "name": "Biodiversite",
            "description": "La biodiversité englobe la variété des formes de vie sur Terre, assurant l''équilibre des écosystèmes et contribuant à la survie de la planète.",
            "quizzes": [
                {
                    "id": 3,
                    "topic": "Biodiversite",
                    "title": "Le Quiz de la Biodiversité Étonnante",
                    "questions": [
                        {
                            "question": "Combien d''espèces animales et végétales sont estimées vivre dans la forêt amazonienne ?",
                            "good_answer": "Des millions",
                            "answers": [
                                "Des milliers",
                                "Des centaines",
                                "Des millions"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande menace pour la biodiversité marine ?",
                            "good_answer": "La pollution plastique",
                            "answers": [
                                "La surpêche",
                                "Le changement climatique",
                                "La pollution plastique"
                            ]
                        },
                        {
                            "question": "Quel est le plus grand animal terrestre de la planète ?",
                            "good_answer": "L''éléphant d''Afrique",
                            "answers": [
                                "Le rhinocéros",
                                "Le lion",
                                "L''éléphant d''Afrique"
                            ]
                        },
                        {
                            "question": "Quelle est la principale cause de la perte de biodiversité dans le monde ?",
                            "good_answer": "La destruction de l''habitat",
                            "answers": [
                                "La chasse excessive",
                                "Les espèces envahissantes",
                                "La destruction de l''habitat"
                            ]
                        },
                        {
                            "question": "Quel oiseau est capable de voler à des altitudes extrêmement élevées et est connu pour ses migrations spectaculaires ?",
                            "good_answer": "L''albatros",
                            "answers": [
                                "Le colibri",
                                "Le moineau",
                                "L''albatros"
                            ]
                        },
                        {
                            "question": "Quelle espèce est souvent considérée comme un symbole de la lutte pour la préservation de la biodiversité ?",
                            "good_answer": "Le panda géant",
                            "answers": [
                                "Le tigre",
                                "Le dauphin",
                                "Le panda géant"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande réserve naturelle du monde, située en Antarctique ?",
                            "good_answer": "La réserve naturelle de l''Antarctique",
                            "answers": [
                                "Le parc national de Yellowstone",
                                "Le parc national des Galápagos",
                                "La réserve naturelle de l''Antarctique"
                            ]
                        },
                        {
                            "question": "Quel est le processus par lequel les espèces évoluent pour s''adapter à leur environnement ?",
                            "good_answer": "La sélection naturelle",
                            "answers": [
                                "La mutation génétique",
                                "La reproduction asexuée",
                                "La sélection naturelle"
                            ]
                        }
                    ]
                },
                {
                    "id": 4,
                    "topic": "Biodiversite",
                    "title": "Quiz sur la Biodiversité",
                    "questions": [
                        {
                            "question": "Combien d''espèces animales sont répertoriées dans le monde ?",
                            "good_answer": "Environ 8,7 millions",
                            "answers": [
                                "Environ 1 million",
                                "Environ 15 millions",
                                "Environ 8,7 millions"
                            ]
                        },
                        {
                            "question": "Quel pourcentage de la biodiversité marine est constitué de poissons ?",
                            "good_answer": "Environ 33%",
                            "answers": [
                                "Environ 10%",
                                "Environ 50%",
                                "Environ 33%"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande menace pour la biodiversité actuellement ?",
                            "good_answer": "La perte d''habitat",
                            "answers": [
                                "La chasse excessive",
                                "La pollution de l''air",
                                "La perte d''habitat"
                            ]
                        },
                        {
                            "question": "Quel est l''objectif principal de la Convention sur la diversité biologique (CDB) ?",
                            "good_answer": "La conservation de la diversité biologique",
                            "answers": [
                                "La promotion de la biotechnologie",
                                "La gestion des déchets",
                                "La conservation de la diversité biologique"
                            ]
                        },
                        {
                            "question": "Quel groupe d''animaux est le plus diversifié en termes d''espèces connues ?",
                            "good_answer": "Les insectes",
                            "answers": [
                                "Les mammifères",
                                "Les reptiles",
                                "Les insectes"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande zone de biodiversité terrestre au monde ?",
                            "good_answer": "La forêt amazonienne",
                            "answers": [
                                "Le désert du Sahara",
                                "La toundra arctique",
                                "La forêt amazonienne"
                            ]
                        },
                        {
                            "question": "Quel est le principal mécanisme de l''évolution de la biodiversité ?",
                            "good_answer": "La sélection naturelle",
                            "answers": [
                                "La migration des espèces",
                                "Les mutations aléatoires",
                                "La sélection naturelle"
                            ]
                        },
                        {
                            "question": "Quel pays abrite la plus grande diversité d''espèces d''oiseaux ?",
                            "good_answer": "La Colombie",
                            "answers": [
                                "Les États-Unis",
                                "L''Australie",
                                "La Colombie"
                            ]
                        }
                    ]
                },
                {
                    "id": 5,
                    "topic": "Biodiversite",
                    "title": "Biodiverse avec moi !",
                    "questions": [
                        {
                            "question": "Quelle est la définition de la biodiversité ?",
                            "good_answer": "La variété des formes de vie sur Terre",
                            "answers": [
                                "La quantité d''eau présente dans les océans",
                                "Le nombre total de pays sur Terre",
                                "La variété des formes de vie sur Terre"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un écosystème ?",
                            "good_answer": "Un ensemble d''organismes vivants et de leur environnement physique",
                            "answers": [
                                "Un type de microscope",
                                "Une substance chimique toxique",
                                "Un ensemble d''organismes vivants et de leur environnement physique"
                            ]
                        },
                        {
                            "question": "Quel est l''objectif de la conservation de la biodiversité ?",
                            "good_answer": "Préserver les espèces et les écosystèmes pour les générations futures",
                            "answers": [
                                "Détruire les habitats naturels",
                                "Promouvoir l''extinction des espèces",
                                "Préserver les espèces et les écosystèmes pour les générations futures"
                            ]
                        },
                        {
                            "question": "Quelle est la principale cause de la perte de biodiversité ?",
                            "good_answer": "La destruction des habitats naturels",
                            "answers": [
                                "Le réchauffement climatique",
                                "La migration des espèces",
                                "La destruction des habitats naturels"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''une espèce endémique ?",
                            "good_answer": "Une espèce présente uniquement dans une région spécifique",
                            "answers": [
                                "Une espèce en voie de disparition",
                                "Une espèce invasive",
                                "Une espèce présente uniquement dans une région spécifique"
                            ]
                        },
                        {
                            "question": "Quels sont les avantages de la biodiversité ?",
                            "good_answer": "La fourniture de ressources alimentaires et médicinales",
                            "answers": [
                                "La pollution de l''environnement",
                                "La destruction des écosystèmes",
                                "La fourniture de ressources alimentaires et médicinales"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un corridor biologique ?",
                            "good_answer": "Une zone qui relie les habitats naturels et permet aux espèces de se déplacer",
                            "answers": [
                                "Un outil utilisé pour mesurer la température",
                                "Un terme désignant un groupe d''oies",
                                "Une zone qui relie les habitats naturels et permet aux espèces de se déplacer"
                            ]
                        },
                        {
                            "question": "Qu''est-ce que la surexploitation des ressources ?",
                            "good_answer": "L''utilisation excessive des ressources naturelles au point de mettre en danger leur survie",
                            "answers": [
                                "La multiplication des ressources naturelles",
                                "La limitation de l''accès aux ressources",
                                "L''utilisation excessive des ressources naturelles au point de mettre en danger leur survie"
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "id": 3,
            "name": "Changement climatique",
            "description" : "Le changement climatique se réfère aux modifications à long terme des conditions météorologiques mondiales, principalement dues à l''activité humaine, avec des impacts sur l''environnement.",
            "quizzes": [
                {
                    "id": 6,
                    "topic": "Changement climatique",
                    "title": "Quiz sur le Changement Climatique",
                    "questions": [
                        {
                            "question": "Quelle est la principale cause du changement climatique actuel ?",
                            "good_answer": "Les émissions de gaz à effet de serre",
                            "answers": [
                                "L''activité volcanique",
                                "Les éruptions solaires",
                                "Les émissions de gaz à effet de serre"
                            ]
                        },
                        {
                            "question": "Quel gaz à effet de serre est le plus préoccupant pour le climat ?",
                            "good_answer": "Le dioxyde de carbone (CO2)",
                            "answers": [
                                "Le méthane (CH4)",
                                "Le protoxyde d''azote (N2O)",
                                "Le dioxyde de carbone (CO2)"
                            ]
                        },
                        {
                            "question": "Quel accord international vise à lutter contre le changement climatique en réduisant les émissions de gaz à effet de serre ?",
                            "good_answer": "L''Accord de Paris",
                            "answers": [
                                "Le Traité de Kyoto",
                                "Le Protocole de Montréal",
                                "L''Accord de Paris"
                            ]
                        },
                        {
                            "question": "Quel phénomène climatique extrême est amplifié par le changement climatique ?",
                            "good_answer": "Les vagues de chaleur",
                            "answers": [
                                "Les blizzards",
                                "Les tornades",
                                "Les vagues de chaleur"
                            ]
                        },
                        {
                            "question": "Quelle est la principale conséquence du réchauffement climatique sur les océans ?",
                            "good_answer": "L''acidification des océans",
                            "answers": [
                                "L''augmentation du niveau de la mer",
                                "La diminution des tempêtes",
                                "L''acidification des océans"
                            ]
                        },
                        {
                            "question": "Quelle région du monde est la plus vulnérable au changement climatique en raison de sa dépendance à l''agriculture ?",
                            "good_answer": "L''Afrique subsaharienne",
                            "answers": [
                                "L''Europe occidentale",
                                "L''Amérique du Nord",
                                "L''Afrique subsaharienne"
                            ]
                        },
                        {
                            "question": "Quel est l''effet du changement climatique sur les glaciers et les calottes glaciaires ?",
                            "good_answer": "La fonte rapide",
                            "answers": [
                                "L''augmentation de la glace",
                                "La stabilisation",
                                "La fonte rapide"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des forêts dans la lutte contre le changement climatique ?",
                            "good_answer": "Elles absorbent le dioxyde de carbone (CO2)",
                            "answers": [
                                "Elles émettent du méthane (CH4)",
                                "Elles n''ont aucun effet sur le climat",
                                "Elles absorbent le dioxyde de carbone (CO2)"
                            ]
                        }
                    ]
                },
                {
                    "id": 7,
                    "topic": "Changement climatique",
                    "title": "Quiz sur le Changement Climatique",
                    "questions": [
                        {
                            "question": "Quelle est la principale cause du changement climatique?",
                            "good_answer": "Les émissions de gaz à effet de serre",
                            "answers": [
                                "Les éruptions volcaniques",
                                "Les rayons cosmiques",
                                "Les émissions de gaz à effet de serre"
                            ]
                        },
                        {
                            "question": "Quel gaz est le plus responsable de l''effet de serre?",
                            "good_answer": "Le dioxyde de carbone (CO2)",
                            "answers": [
                                "L''oxygène (O2)",
                                "L''azote (N2)",
                                "Le dioxyde de carbone (CO2)"
                            ]
                        },
                        {
                            "question": "Qu''est-ce que l''effet de serre?",
                            "good_answer": "Le phénomène par lequel certaines substances emprisonnent la chaleur dans l''atmosphère terrestre.",
                            "answers": [
                                "Un vent fort",
                                "Un courant océanique",
                                "Le phénomène par lequel certaines substances emprisonnent la chaleur dans l''atmosphère terrestre."
                            ]
                        },
                        {
                            "question": "Quelle est la conséquence du réchauffement climatique sur les glaciers?",
                            "good_answer": "La fonte des glaciers",
                            "answers": [
                                "La croissance des glaciers",
                                "La couleur des glaciers devient rouge",
                                "La fonte des glaciers"
                            ]
                        },
                        {
                            "question": "Quel est le principal secteur émetteur de gaz à effet de serre?",
                            "good_answer": "Le secteur de l''énergie",
                            "answers": [
                                "L''agriculture",
                                "Le secteur de la santé",
                                "Le secteur de l''énergie"
                            ]
                        },
                        {
                            "question": "Quel accord international vise à lutter contre le changement climatique?",
                            "good_answer": "L''Accord de Paris",
                            "answers": [
                                "L''Accord de Rome",
                                "L''Accord de Londres",
                                "L''Accord de Paris"
                            ]
                        },
                        {
                            "question": "Quelle est la conséquence du changement climatique sur les océans?",
                            "good_answer": "L''acidification des océans",
                            "answers": [
                                "Augmentation de la salinité des océans",
                                "Diminution du niveau de la mer",
                                "L''acidification des océans"
                            ]
                        },
                        {
                            "question": "Quel est le rôle de la forêt dans la lutte contre le changement climatique?",
                            "good_answer": "Absorber le dioxyde de carbone de l''atmosphère",
                            "answers": [
                                "Libérer du dioxyde de carbone dans l''atmosphère",
                                "Détruire la couche d''ozone",
                                "Absorber le dioxyde de carbone de l''atmosphère"
                            ]
                        }
                    ]
                },
                {
                    "id": 8,
                    "topic": "Changement climatique",
                    "title": "Quiz sur le Réchauffement Climatique",
                    "questions": [
                        {
                            "question": "Quel est l''impact du réchauffement climatique sur la fréquence des phénomènes météorologiques extrêmes ?",
                            "good_answer": "Ils deviennent plus fréquents",
                            "answers": [
                                "Ils deviennent moins fréquents",
                                "Ils restent inchangés",
                                "Ils deviennent plus fréquents"
                            ]
                        },
                        {
                            "question": "Quelle est la principale source d''émission de gaz à effet de serre liée aux activités humaines ?",
                            "good_answer": "La combustion des énergies fossiles",
                            "answers": [
                                "L''agriculture",
                                "La déforestation",
                                "La combustion des énergies fossiles"
                            ]
                        },
                        {
                            "question": "Quel phénomène climatique est associé au réchauffement des océans ?",
                            "good_answer": "L''élévation du niveau de la mer",
                            "answers": [
                                "La diminution des tempêtes",
                                "La formation d''icebergs",
                                "L''élévation du niveau de la mer"
                            ]
                        },
                        {
                            "question": "Quel est le principal contributeur au trou dans la couche d''ozone, un problème lié au changement climatique ?",
                            "good_answer": "Les gaz réfrigérants",
                            "answers": [
                                "Les émissions de CO2",
                                "Les poussières atmosphériques",
                                "Les gaz réfrigérants"
                            ]
                        },
                        {
                            "question": "Quel est l''effet du réchauffement climatique sur la biodiversité ?",
                            "good_answer": "La perte de diversité biologique",
                            "answers": [
                                "L''augmentation des espèces",
                                "La migration des espèces",
                                "La perte de diversité biologique"
                            ]
                        },
                        {
                            "question": "Quelle est la principale conséquence du réchauffement climatique sur les régions polaires ?",
                            "good_answer": "La fonte accélérée des glaciers",
                            "answers": [
                                "L''augmentation de la banquise",
                                "La stabilisation des températures",
                                "La fonte accélérée des glaciers"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des océans dans l''absorption du surplus de chaleur dû au réchauffement climatique ?",
                            "good_answer": "Ils agissent comme un réservoir thermique",
                            "answers": [
                                "Ils amplifient le réchauffement",
                                "Ils sont insensibles à la chaleur",
                                "Ils agissent comme un réservoir thermique"
                            ]
                        },
                        {
                            "question": "Quel gaz est libéré par la fonte du permafrost, contribuant au réchauffement climatique ?",
                            "good_answer": "Le méthane",
                            "answers": [
                                "Le dioxyde de carbone",
                                "L''oxygène",
                                "Le méthane"
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "id": 4,
            "name": "Ecologie Marine",
            "description" : "L''écologie marine étudie les interactions entre les organismes marins et leur environnement, crucial pour la compréhension et la préservation des océans.",
            "quizzes": [
                {
                    "id": 9,
                    "topic": "Ecologie Marine",
                    "title": "Quiz sur l''écologie marine",
                    "questions": [
                        {
                            "question": "Quelle est la principale menace pour les récifs coralliens ?",
                            "good_answer": "Le réchauffement climatique",
                            "answers": [
                                "La pollution plastique",
                                "La surpêche",
                                "Le réchauffement climatique"
                            ]
                        },
                        {
                            "question": "Quel est le plus grand animal de tous les temps ?",
                            "good_answer": "La baleine bleue",
                            "answers": [
                                "Le requin blanc",
                                "L''orque",
                                "La baleine bleue"
                            ]
                        },
                        {
                            "question": "Comment s''appelle la zone en pleine mer où la vie est très abondante ?",
                            "good_answer": "La zone pélagique",
                            "answers": [
                                "La zone benthique",
                                "La zone abyssale",
                                "La zone pélagique"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un herbier marin ?",
                            "good_answer": "Une prairie sous-marine d''herbes marines",
                            "answers": [
                                "Un corail géant",
                                "Un poisson de grande taille",
                                "Une prairie sous-marine d''herbes marines"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un écosystème côtier ?",
                            "good_answer": "Un ensemble de communautés vivantes entre la terre et la mer",
                            "answers": [
                                "Un écosystème situé au fond de l''océan",
                                "Un espace sans vie marine",
                                "Un ensemble de communautés vivantes entre la terre et la mer"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un tsunami ?",
                            "good_answer": "Une série de vagues provoquée par un séisme sous-marin",
                            "answers": [
                                "Un mouvement régulier des marées",
                                "Un tourbillon marin",
                                "Une série de vagues provoquée par un séisme sous-marin"
                            ]
                        },
                        {
                            "question": "Que sont les espèces invasives marines ?",
                            "good_answer": "Des espèces non indigènes qui envahissent un écosystème marin",
                            "answers": [
                                "Des espèces endémiques protégées",
                                "Des espèces rares et menacées",
                                "Des espèces non indigènes qui envahissent un écosystème marin"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un récif artificiel ?",
                            "good_answer": "Une structure créée par l''homme pour favoriser la biodiversité marine",
                            "answers": [
                                "Un récif naturel formé par des coraux",
                                "Un récif sans vie marine",
                                "Une structure créée par l''homme pour favoriser la biodiversité marine"
                            ]
                        }
                    ]
                },
                {
                    "id": 10,
                    "topic": "Écologie marine",
                    "title": "Quiz sur l''Écologie Marine",
                    "questions": [
                        {
                            "question": "Quel est le plus grand contributeur à la pollution plastique des océans ?",
                            "good_answer": "Les déchets plastiques à usage unique",
                            "answers": [
                                "Les emballages alimentaires en carton",
                                "Les déchets métalliques",
                                "Les déchets plastiques à usage unique"
                            ]
                        },
                        {
                            "question": "Quel est le principal danger pour les récifs coralliens dans le contexte de l''écologie marine ?",
                            "good_answer": "Le blanchissement corallien",
                            "answers": [
                                "L''acidification des océans",
                                "La surpêche",
                                "Le blanchissement corallien"
                            ]
                        },
                        {
                            "question": "Quelle est la principale menace pour les tortues marines en matière d''écologie marine ?",
                            "good_answer": "Les filets de pêche abandonnés",
                            "answers": [
                                "Le changement climatique",
                                "La destruction de l''habitat",
                                "Les filets de pêche abandonnés"
                            ]
                        },
                        {
                            "question": "Quel est l''effet de la marée noire sur l''écosystème marin ?",
                            "good_answer": "La contamination des espèces marines",
                            "answers": [
                                "La stimulation de la croissance des algues",
                                "L''amélioration de la biodiversité",
                                "La contamination des espèces marines"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des mangroves dans l''équilibre de l''écosystème marin ?",
                            "good_answer": "Elles servent de nurserie pour de nombreuses espèces marines",
                            "answers": [
                                "Elles sont des lieux de reproduction pour les méduses",
                                "Elles n''ont aucun impact sur l''écologie marine",
                                "Elles servent de nurserie pour de nombreuses espèces marines"
                            ]
                        },
                        {
                            "question": "Quelle est la principale source de surpêche dans le monde, menaçant l''écologie marine ?",
                            "good_answer": "La pêche illégale, non déclarée et non réglementée (INN)",
                            "answers": [
                                "La pêche artisanale",
                                "La pêche industrielle responsable",
                                "La pêche illégale, non déclarée et non réglementée (INN)"
                            ]
                        },
                        {
                            "question": "Quel est l''impact de l''acidification des océans sur les coquillages et les mollusques ?",
                            "good_answer": "Elle affaiblit leurs coquilles",
                            "answers": [
                                "Elle renforce leurs coquilles",
                                "Elle n''affecte pas les coquillages",
                                "Elle affaiblit leurs coquilles"
                            ]
                        },
                        {
                            "question": "Quelle est la conséquence de la destruction des herbiers marins sur l''écologie marine ?",
                            "good_answer": "La perte d''habitats essentiels pour de nombreuses espèces",
                            "answers": [
                                "L''amélioration de la qualité de l''eau",
                                "L''augmentation de la productivité marine",
                                "La perte d''habitats essentiels pour de nombreuses espèces"
                            ]
                        }
                    ]
                },

                {
                    "id": 11,
                    "topic": "Écologie marine",
                    "title": "Quiz sur l''Écologie Marine",
                    "questions": [
                        {
                            "question": "Quel est le rôle des herbiers marins dans l''écosystème marin ?",
                            "good_answer": "Ils servent de nurserie et d''habitat pour de nombreuses espèces",
                            "answers": [
                                "Ils produisent de l''oxygène pour les poissons",
                                "Ils n''ont aucun impact sur l''écosystème marin",
                                "Ils servent de nurserie et d''habitat pour de nombreuses espèces"
                            ]
                        },
                        {
                            "question": "Quelle est la principale menace pour les requins dans le contexte de l''écologie marine ?",
                            "good_answer": "La surpêche et le commerce des ailerons de requin",
                            "answers": [
                                "La perte d''habitats",
                                "Les changements climatiques",
                                "La surpêche et le commerce des ailerons de requin"
                            ]
                        },
                        {
                            "question": "Quel est l''effet des marées rouges sur la faune marine ?",
                            "good_answer": "Elles peuvent causer des mortalités massives de poissons et de coquillages",
                            "answers": [
                                "Elles favorisent la croissance des algues",
                                "Elles n''ont aucun impact sur la faune marine",
                                "Elles peuvent causer des mortalités massives de poissons et de coquillages"
                            ]
                        },
                        {
                            "question": "Quel est l''impact de la pollution sonore sur les cétacés, tels que les baleines ?",
                            "good_answer": "Elle perturbe leur communication et leur orientation",
                            "answers": [
                                "Elle renforce leur capacité d''écholocation",
                                "Elle n''affecte pas les cétacés",
                                "Elle perturbe leur communication et leur orientation"
                            ]
                        },
                        {
                            "question": "Quelle est la conséquence de la pollution par les hydrocarbures sur les oiseaux marins ?",
                            "good_answer": "Elle peut causer des dommages aux plumes et à la structure des plumes",
                            "answers": [
                                "Elle améliore la flottabilité des oiseaux",
                                "Elle n''a aucun effet sur les oiseaux marins",
                                "Elle peut causer des dommages aux plumes et à la structure des plumes"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des récifs coralliens dans la protection des côtes contre les tempêtes ?",
                            "good_answer": "Ils agissent comme des barrières naturelles",
                            "answers": [
                                "Ils amplifient l''impact des tempêtes",
                                "Ils n''ont aucun effet sur les tempêtes",
                                "Ils agissent comme des barrières naturelles"
                            ]
                        },
                        {
                            "question": "Quel est l''effet de la montée du niveau de la mer sur les écosystèmes côtiers ?",
                            "good_answer": "Elle entraîne la perte d''habitats côtiers",
                            "answers": [
                                "Elle favorise la biodiversité côtière",
                                "Elle n''a aucun impact sur les écosystèmes côtiers",
                                "Elle entraîne la perte d''habitats côtiers"
                            ]
                        },
                        {
                            "question": "Quel est le principal problème lié à la pêche fantôme dans l''écologie marine ?",
                            "good_answer": "Les engins de pêche abandonnés qui continuent de capturer des animaux marins",
                            "answers": [
                                "La pêche excessive",
                                "La destruction des habitats marins",
                                "Les engins de pêche abandonnés qui continuent de capturer des animaux marins"
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "id": 5,
            "name": "Protection des zones humides",
            "description" : "La protection des zones humides vise à préserver ces écosystèmes essentiels qui fournissent de l''eau douce, régulent les inondations et abritent une biodiversité unique.",
            "quizzes": [
                {
                    "id": 12,
                    "topic": "Protection des zones humides",
                    "quiz": {
                        "title": "Quiz de l''Envol du Héron",
                        "questions": [
                            {
                                "question": "Quelle est la principale menace pour les zones humides ?",
                                "good_answer": "Le drainage excessif",
                                "answers": [
                                    "Les castors",
                                    "Les chauves-souris",
                                    "Le drainage excessif"
                                ]
                            },
                            {
                                "question": "Quel est le rôle des zones humides dans l''écosystème ?",
                                "good_answer": "Elles filtrent les eaux et purifient l''air",
                                "answers": [
                                    "Elles sont des habitats pour les dromadaires",
                                    "Elles sont des plateformes de lancement de fusées",
                                    "Elles filtrent les eaux et purifient l''air"
                                ]
                            },
                            {
                                "question": "Comment peut-on protéger les zones humides ?",
                                "good_answer": "En limitant les activités humaines impactantes",
                                "answers": [
                                    "En construisant des gratte-ciels en leur centre",
                                    "En organisant des courses de bateaux à moteur",
                                    "En limitant les activités humaines impactantes"
                                ]
                            },
                            {
                                "question": "Quel oiseau est emblématique des zones humides ?",
                                "good_answer": "Le héron cendré",
                                "answers": [
                                    "Le manchot empereur",
                                    "Le pigeon voyageur",
                                    "Le héron cendré"
                                ]
                            },
                            {
                                "question": "Quel est l''effet des zones humides sur les inondations ?",
                                "good_answer": "Elles les atténuent en absorbant l''eau en excès",
                                "answers": [
                                    "Elles les aggravent en utilisant des pompiers",
                                    "Elles les ignorent et font des châteaux de sable",
                                    "Elles les atténuent en absorbant l''eau en excès"
                                ]
                            },
                            {
                                "question": "Qu''est-ce qu''un marais ?",
                                "good_answer": "Un écosystème humide avec une végétation aquatique flottante",
                                "answers": [
                                    "Une tartine beurrée",
                                    "Une danse traditionnelle",
                                    "Un écosystème humide avec une végétation aquatique flottante"
                                ]
                            },
                            {
                                "question": "Pourquoi est-il important de préserver les zones humides ?",
                                "good_answer": "Elles abritent une biodiversité exceptionnelle",
                                "answers": [
                                    "Elles sont le repaire des extraterrestres",
                                    "Elles ont des restaurants étoilés Michelin",
                                    "Elles abritent une biodiversité exceptionnelle"
                                ]
                            },
                            {
                                "question": "Quelle est la meilleure façon de visiter une zone humide ?",
                                "good_answer": "En respectant la réglementation en vigueur",
                                "answers": [
                                    "En organisant une rave-party géante",
                                    "En portant un sombrero fluorescent",
                                    "En respectant la réglementation en vigueur"
                                ]
                            }
                        ]
                    }
                },
                {
                    "id": 13,
                    "topic": "Protection des zones humides",
                    "quiz": {
                        "title": "Quiz rigolo sur la protection des zones humides",
                        "questions": [
                            {
                                "question": "Quelle est la principale cause de dégradation des zones humides ?",
                                "good_answer": "La pollution",
                                "answers": [
                                    "Les oiseaux",
                                    "La danse de la pluie",
                                    "La pollution"
                                ]
                            },
                            {
                                "question": "Quel est l''animal emblématique des zones humides ?",
                                "good_answer": "Le castor",
                                "answers": [
                                    "La licorne",
                                    "Le kangourou",
                                    "Le castor"
                                ]
                            },
                            {
                                "question": "Qu''est-ce qu''un marais ?",
                                "good_answer": "Une zone humide principalement composée d''herbes et de roseaux",
                                "answers": [
                                    "Un dessert sucré",
                                    "Un type de chapeau",
                                    "Une zone humide principalement composée d''herbes et de roseaux"
                                ]
                            },
                            {
                                "question": "Quel est l''effet bénéfique des zones humides sur l''environnement ?",
                                "good_answer": "Elles filtrent et purifient l''eau",
                                "answers": [
                                    "Elles rendent les grenouilles sympathiques",
                                    "Elles produisent des arcs-en-ciel",
                                    "Elles filtrent et purifient l''eau"
                                ]
                            },
                            {
                                "question": "Quelle est la plus grande zone humide du monde ?",
                                "good_answer": "Le Pantanal au Brésil",
                                "answers": [
                                    "Mon jardin",
                                    "Le désert du Sahara",
                                    "Le Pantanal au Brésil"
                                ]
                            },
                            {
                                "question": "Comment appelle-t-on un petit étang formé dans une zone humide ?",
                                "good_answer": "Une mare",
                                "answers": [
                                    "Un bouillon de culture",
                                    "Un jacuzzi",
                                    "Une mare"
                                ]
                            },
                            {
                                "question": "Quelle est la couleur de la vase présente dans les zones humides ?",
                                "good_answer": "Marron",
                                "answers": [
                                    "Orange fluo",
                                    "Arc-en-ciel",
                                    "Marron"
                                ]
                            },
                            {
                                "question": "Qu''est-ce qu''une tourbière ?",
                                "good_answer": "Une zone humide où se forme progressivement une sorte de terre spéciale appelée tourbe",
                                "answers": [
                                    "Une tour en béton",
                                    "Une espèce de biscuit",
                                    "Une zone humide où se forme progressivement une sorte de terre spéciale appelée tourbe"
                                ]
                            }
                        ]
                    }
                },
                {
                    "id": 14,
                    "topic": "Protection des zones humides",
                    "quiz": {
                        "title": "Le marécage sophistiqué",
                        "questions": [
                            {
                                "question": "Quel pourcentage des zones humides mondiales ont disparu au cours des 100 dernières années ?",
                                "good_answer": "50%",
                                "answers": [
                                    "5%",
                                    "25%",
                                    "50%"
                                ]
                            },
                            {
                                "question": "Quelle est la plus grande zone humide du monde en termes de superficie ?",
                                "good_answer": "Le Pantanal",
                                "answers": [
                                    "Le Marais Poitevin",
                                    "Les Everglades",
                                    "Le Pantanal"
                                ]
                            },
                            {
                                "question": "Combien d''espèces d''oiseaux dépendent des zones humides pour leur survie ?",
                                "good_answer": "Plus de 1 000",
                                "answers": [
                                    "Moins de 100",
                                    "Environ 500",
                                    "Plus de 1 000"
                                ]
                            },
                            {
                                "question": "Quelle proportion des espèces menacées d''extinction vivent dans les zones humides ?",
                                "good_answer": "Environ 40%",
                                "answers": [
                                    "Moins de 10%",
                                    "Environ 20%",
                                    "Environ 40%"
                                ]
                            },
                            {
                                "question": "Quelle est la principale cause de la destruction des zones humides dans le monde ?",
                                "good_answer": "L''urbanisation",
                                "answers": [
                                    "Le réchauffement climatique",
                                    "L''agriculture intensive",
                                    "L''urbanisation"
                                ]
                            }
                        ]
                    }
                }
            ]
        },
        {
            "id": 6,
            "name": "Conservation de l''eau",
            "description" : "La conservation de l''eau implique des pratiques visant à économiser et à gérer judicieusement cette ressource précieuse pour répondre aux besoins actuels et futurs.",
            "quizzes": [
                {
                    "id": 15,
                    "topic": "Conservation de l''eau",
                    "quiz": {
                        "title": "Le joueur d''eau",
                        "questions": [
                            {
                                "question": "Quelle est la quantité d''eau nécessaire pour qu''un humain survive pendant une semaine sans aucune autre source de nourriture ?",
                                "good_answer": "7 litres",
                                "answers": [
                                    "1 litre",
                                    "20 litres",
                                    "7 litres"
                                ]
                            },
                            {
                                "question": "Quelle est la principale cause de perte d''eau potable dans les fuites d''un système de distribution d''eau ?",
                                "good_answer": "Les canalisations défectueuses",
                                "answers": [
                                    "Les robinets qui restent ouverts",
                                    "L''évaporation naturelle",
                                    "Les canalisations défectueuses"
                                ]
                            },
                            {
                                "question": "Quelle partie du monde est le plus touchée par le stress hydrique ?",
                                "good_answer": "Le Moyen-Orient",
                                "answers": [
                                    "L''Amérique du Nord",
                                    "L''Australie",
                                    "Le Moyen-Orient"
                                ]
                            },
                            {
                                "question": "Quelle est la proportion d''eau salée sur Terre ?",
                                "good_answer": "97 %",
                                "answers": [
                                    "50 %",
                                    "75 %",
                                    "97 %"
                                ]
                            },
                            {
                                "question": "Dans quelle activité quotidienne consommons-nous le plus d''eau en moyenne ?",
                                "good_answer": "L''agriculture",
                                "answers": [
                                    "La consommation domestique",
                                    "L''industrie",
                                    "L''agriculture"
                                ]
                            }
                        ]
                    }
                },
                {
                    "id": 16,
                    "topic": "Conservation de l''eau",
                    "quiz": {
                        "title": "Plouf Plouf à la rescousse",
                        "questions": [
                            {
                                "question": "Quelle est la quantité d''eau douce disponible sur Terre ?",
                                "good_answer": "2,5%",
                                "answers": [
                                    "50%",
                                    "85%",
                                    "2,5%"
                                ]
                            },
                            {
                                "question": "Quelle est la consommation moyenne d''eau d''une chasse d''eau traditionnelle ?",
                                "good_answer": "9 à 12 litres",
                                "answers": [
                                    "2 à 4 litres",
                                    "15 à 18 litres",
                                    "9 à 12 litres"
                                ]
                            },
                            {
                                "question": "Quelle activité humaine utilise le plus d''eau dans le monde ?",
                                "good_answer": "L''agriculture",
                                "answers": [
                                    "La production d''électricité",
                                    "La consommation domestique",
                                    "L''agriculture"
                                ]
                            },
                            {
                                "question": "Combien de litres d''eau faut-il pour produire une seule canette de soda ?",
                                "good_answer": "250 à 300 litres",
                                "answers": [
                                    "50 à 100 litres",
                                    "500 à 600 litres",
                                    "250 à 300 litres"
                                ]
                            },
                            {
                                "question": "Quel est le pourcentage d''eau froide dans un iceberg ?",
                                "good_answer": "90%",
                                "answers": [
                                    "50%",
                                    "25%",
                                    "90%"
                                ]
                            },
                            {
                                "question": "Quel pays possède le plus grand pourcentage de sa population sans accès à l''eau potable ?",
                                "good_answer": "Somalie",
                                "answers": [
                                    "Inde",
                                    "Brésil",
                                    "Somalie"
                                ]
                            },
                            {
                                "question": "Combien de litres d''eau faut-il pour produire une tonne de papier recyclé ?",
                                "good_answer": "250 000 litres",
                                "answers": [
                                    "50 000 litres",
                                    "500 000 litres",
                                    "250 000 litres"
                                ]
                            }
                        ]
                    }
                },
                {
                    "id": 17,
                    "topic": "Conservation de l''eau",
                    "quiz": {
                        "title": "Pourquoi l''eau ne dit jamais de blagues ?",
                        "questions": [
                            {
                                "question": "Quelle est la principale source d''eau douce utilisée par les humains dans le monde ?",
                                "good_answer": "Les eaux souterraines",
                                "answers": [
                                    "Les rivières",
                                    "Les lacs",
                                    "Les glaciers",
                                    "Les eaux souterraines"
                                ]
                            },
                            {
                                "question": "Combien de litres d''eau utilise-t-on en moyenne par jour pour une douche de 10 minutes ?",
                                "good_answer": "Environ 80 litres",
                                "answers": [
                                    "Environ 20 litres",
                                    "Environ 50 litres",
                                    "Environ 120 litres",
                                    "Environ 80 litres"
                                ]
                            },
                            {
                                "question": "Qu''est-ce que la xéropaysagisme ?",
                                "good_answer": "Une technique de jardinage permettant de réduire la consommation d''eau",
                                "answers": [
                                    "Une maladie des plantes causée par un excès d''arrosage",
                                    "La construction d''un barrage pour stocker de l''eau",
                                    "Une méthode pour purifier l''eau",
                                    "Une technique de jardinage permettant de réduire la consommation d''eau"
                                ]
                            }
                        ]
                    }
                },
                {
                    "id": 18,
                    "topic": "Conservation de l''eau",
                    "quiz": {
                        "title": "Quand une goutte d''eau se transforme en super-héros",
                        "questions": [
                            {
                                "question": "Quelle est la principale cause de la pénurie d''eau dans de nombreuses régions du monde ?",
                                "good_answer": "Le gaspillage excessif d''eau",
                                "answers": [
                                    "La pollution de l''eau",
                                    "Les précipitations insuffisantes",
                                    "La surconsommation de boissons gazeuses",
                                    "Le gaspillage excessif d''eau"
                                ]
                            },
                            {
                                "question": "Quel pourcentage de l''eau sur Terre est disponible pour la consommation humaine ?",
                                "good_answer": "Moins de 1%",
                                "answers": [
                                    "Environ 25%",
                                    "Près de 50%",
                                    "Plus de 75%",
                                    "Moins de 1%"
                                ]
                            },
                            {
                                "question": "Qu''est-ce que l''agriculture durable peut contribuer à la conservation de l''eau ?",
                                "good_answer": "La mise en place de techniques d''irrigation efficaces",
                                "answers": [
                                    "L''abandon de l''agriculture",
                                    "L''utilisation intensive d''engrais chimiques",
                                    "L''augmentation de la consommation d''eau dans les exploitations agricoles",
                                    "La mise en place de techniques d''irrigation efficaces"
                                ]
                            }
                        ]
                    }
                },
                {
                    "id": 19,
                    "topic": "Conservation de l''eau",
                    "quiz": {
                        "title": "Axoloto : Pokemon de type Eau",
                        "questions": [
                            {
                                "question": "Quelle est la meilleure façon de conserver l''eau dans une maison ?",
                                "good_answer": "Installer des pommeaux de douche à faible débit",
                                "answers": [
                                    "Laisser le robinet ouvert pendant que vous vous brossez les dents",
                                    "Remplir la baignoire à ras bord pour chaque bain",
                                    "Arroser le jardin tous les jours",
                                    "Installer des pommeaux de douche à faible débit"
                                ]
                            },
                            {
                                "question": "Combien d''eau est gaspillée par une fuite de robinet qui goutte constamment, chaque jour ?",
                                "good_answer": "Environ 136 litres",
                                "answers": [
                                    "Environ 5 litres",
                                    "Environ 20 litres",
                                    "Aucune eau n''est gaspillée",
                                    "Environ 136 litres"
                                ]
                            },
                            {
                                "question": "Qu''est-ce que la méthode de l''irrigation goutte à goutte ?",
                                "good_answer": "Une méthode qui utilise des tuyaux à bas débit pour apporter de petites quantités d''eau directement aux plantes",
                                "answers": [
                                    "Une méthode qui utilise des arroseurs à haute pression pour mouiller toute la zone du jardin",
                                    "Une méthode qui n''utilise pas du tout d''eau pour l''arrosage des plantes",
                                    "Une méthode qui implique de remplir une cuve d''eau et d''arroser manuellement les plantes",
                                    "Une méthode qui utilise des tuyaux à bas débit pour apporter de petites quantités d''eau directement aux plantes"
                                ]
                            }
                        ]
                    }
                },
                {
                    "id": 20,
                    "topic": "Conservation de l''eau",
                    "quiz": {
                        "questions": [
                            {
                                "question": "Quelle est la principale raison de la conservation de l''eau ?",
                                "good_answer": "Préserver les ressources en eau pour les générations futures",
                                "answers": [
                                    "Réduire les coûts de factures d''eau",
                                    "Éviter les restrictions d''eau imposées par le gouvernement",
                                    "Maintenir les installations de traitement de l''eau en bon état",
                                    "Préserver les ressources en eau pour les générations futures"
                                ]
                            },
                            {
                                "question": "Combien de litres d''eau peut être économisé en utilisant un pommeau de douche à faible débit par minute ?",
                                "good_answer": "Environ 9 à 11 litres",
                                "answers": [
                                    "Environ 3 à 5 litres",
                                    "Environ 15 à 20 litres",
                                    "Environ 25 à 30 litres",
                                    "Environ 9 à 11 litres"
                                ]
                            },
                            {
                                "question": "Qu''est-ce que l''irrigation goutte-à-goutte dans le contexte de la conservation de l''eau ?",
                                "good_answer": "Une méthode d''arrosage qui fournit de l''eau directement aux racines des plantes",
                                "answers": [
                                    "Une technique d''arrosage qui nécessite une grande quantité d''eau",
                                    "Une pratique qui n''est pas adaptée aux climats arides",
                                    "Une méthode qui gaspille plus d''eau que d''autres systèmes d''irrigation",
                                    "Une méthode d''arrosage qui fournit de l''eau directement aux racines des plantes"
                                ]
                            }
                        ]
                    }
                }
            ]
        }
    ],
    "version": 1
}'
           , NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM quiz_games
WHERE gid = (SELECT id FROM games WHERE name = 'Quiz' AND game_version = '1.0');
-- +goose StatementEnd
