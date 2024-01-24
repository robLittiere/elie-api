-- +goose Up
-- +goose StatementBegin
INSERT INTO quiz_games (gid, data, created_at, updated_at)
values ((SELECT id FROM games WHERE name = 'Quiz' AND game_version = '1.0'),
           '{
    "topic": [
        {
            "id": 1,
            "name": "Energie solaire",
            "description" : "L''énergie solaire produit de l''électricité propre et renouvelable en exploitant la puissance du soleil.",
            "quizzes": [
                {
                    "id": 1,
                    "topic": "Energie solaire",
                    "title": "Qu''est-ce-que l''énergie solaire ?",
                    "questions": [
                        {
                            "question": "Quel est le processus par lequel l''énergie solaire est convertie en électricité ?",
                            "good_answer": "La conversion photovoltaïque",
                            "answers": [
                                "La photosynthèse",
                                "La conversion photovoltaïque",
                                "La combustion solaire",
                                "La capture radiante"
                            ]
                        },
                        {
                            "question": "Quel élément est principalement utilisé dans les panneaux solaires pour capter la lumière du soleil ?",
                            "good_answer": "Le silicium",
                            "answers": [
                                "L''hydrogène",
                                "Le carbone",
                                "Le silicium",
                                "Le titane"
                            ]
                        },
                        {
                            "question": "Quel est l''appareil qui permet de suivre la trajectoire du soleil pour maximiser la capture d''énergie solaire ?",
                            "good_answer": "Le suiveur solaire",
                            "answers": [
                                "Le traqueur solaire",
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
                                "Joule-solaire (Js)",
                                "Watt-crête (Wc)"
                            ]
                        },
                        {
                            "question": "Quel pays est le plus grand producteur d''énergie solaire au monde ?",
                            "good_answer": "La Chine",
                            "answers": [
                                "Les États-Unis",
                                "La Chine",
                                "L''Inde",
                                "La France"
                            ]
                        },
                        {
                            "question": "Quel avantage majeur de l''énergie solaire en fait une source d''énergie renouvelable attrayante ?",
                            "good_answer": "Elle ne produit pas de pollution atmosphérique",
                            "answers": [
                                "Elle ne produit pas de pollution atmosphérique",
                                "Elle est bon marché",
                                "Elle est facile à stocker",
                                "Elle nécessite une faible consommation d''eau pour sa production"
                            ]
                        },
                        {
                            "question": "Quel dispositif permet de stocker l''énergie solaire pour une utilisation ultérieure ?",
                            "good_answer": "Les batteries solaires",
                            "answers": [
                                "Les réservoirs à rayons solaires",
                                "Les panneaux solaires",
                                "Les onduleurs solaires",
                                "Les batteries solaires"
                            ]
                        },
                        {
                            "question": "Quelle est l''efficacité typique des panneaux solaires commerciaux ?",
                            "good_answer": "Entre 15% et 20%",
                            "answers": [
                                "Moins de 15%",
                                "Entre 15% et 20%",
                                "Entre 25% et 50%",
                                "Plus de 50%"
                            ]
                        }
                    ]
                },
                {
                    "id": 2,
                    "topic": "Energie solaire",
                    "title": "Comment fonctionne l''énergie solaire ?",
                    "questions": [
                        {
                            "question": "Quelle est la source d''énergie principale du solaire photovoltaïque ?",
                            "good_answer": "Lumière du soleil",
                            "answers": [
                                "Vent",
                                "Lumière du soleil",
                                "Pétrole",
                                "Pluie"
                            ]
                        },
                        {
                            "question": "Quel composant clé est utilisé pour convertir la lumière en électricité dans les panneaux solaires ?",
                            "good_answer": "Cellules photovoltaïques",
                            "answers": [
                                "Batteries",
                                "Aimants",
                                "Bobines électroluminescentes",
                                "Cellules photovoltaïques"
                            ]
                        },
                        {
                            "question": "Quel pays est le plus grand producteur d''énergie solaire au monde ?",
                            "good_answer": "Chine",
                            "answers": [
                                "Chine",
                                "États-Unis",
                                "Canada",
                                "Inde"
                            ]
                        },
                        {
                            "question": "Comment est stockée l''énergie solaire pour une utilisation ultérieure ?",
                            "good_answer": "Batteries",
                            "answers": [
                                "Réservoirs d''eau",
                                "Sacs de sable",
                                "Batteries",
                                "Cuves à photons"
                            ]
                        },
                        {
                            "question": "Quel est l''angle optimal pour l''installation de panneaux solaires afin de maximiser l''exposition au soleil ?",
                            "good_answer": "Entre 30 et 45 degrés",
                            "answers": [
                                "0 degrés",
                                "Entre 30 et 45 degrés",
                                "Entre 60 et 75 degrés",
                                "90 degrés"
                            ]
                        },
                        {
                            "question": "Quel type de rayonnement solaire est converti en énergie électrique par les panneaux solaires ?",
                            "good_answer": "Rayonnement solaire photovoltaïque",
                            "answers": [
                                "Rayonnement gamma solaire",
                                "Rayonnement infrarouge",
                                "Rayonnement ultraviolet",
                                "Rayonnement solaire photovoltaïque"
                            ]
                        },
                        {
                            "question": "Quelle est la durée de vie typique des panneaux solaires ?",
                            "good_answer": "25 ans",
                            "answers": [
                                "5 ans",
                                "15 ans",
                                "25 ans",
                                "35 ans"
                            ]
                        },
                        {
                            "question": "Quelle est la principale source d''énergie utilisée pour chauffer l''eau à l''aide de l''énergie solaire ?",
                            "good_answer": "Énergie solaire thermique",
                            "answers": [
                                "Gaz naturel",
                                "Électricité",
                                "Énergie hydrothermique",
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
            "description": "La biodiversité assure l''équilibre des écosystèmes et la survie de la planète.",
            "quizzes": [
                {
                    "id": 3,
                    "topic": "Biodiversite",
                    "title": "Qu''est-ce-que la biodiversité ?",
                    "questions": [
                        {
                            "question": "Combien d''espèces animales et végétales sont estimées vivre dans la forêt amazonienne ?",
                            "good_answer": "Des millions",
                            "answers": [
                                "Des centaines",
                                "Des milliers",
                                "Des millions",
                                "Des milliards"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande menace pour la biodiversité marine ?",
                            "good_answer": "La pollution plastique",
                            "answers": [
                                "La surpêche",
                                "L''acidification des océans",
                                "Le changement climatique",
                                "La pollution plastique"
                            ]
                        },
                        {
                            "question": "Quel est le plus grand animal terrestre de la planète ?",
                            "good_answer": "L''éléphant d''Afrique",
                            "answers": [
                                "L''éléphant d''Afrique",
                                "Le rhinocéros",
                                "Le lion",
                                "Le cheval"
                            ]
                        },
                        {
                            "question": "Quelle est la principale cause de la perte de biodiversité dans le monde ?",
                            "good_answer": "La destruction de l''habitat",
                            "answers": [
                                "La chasse excessive",
                                "La destruction de l''habitat",
                                "Les espèces envahissantes",
                                "Le réchauffement climatique"
                            ]
                        },
                        {
                            "question": "Quel oiseau est capable de voler à des altitudes extrêmement élevées et est connu pour ses migrations spectaculaires ?",
                            "good_answer": "L''albatros",
                            "answers": [
                                "Le colibri",
                                "Le moineau",
                                "L''aigle",
                                "L''albatros"
                            ]
                        },
                        {
                            "question": "Quelle espèce est souvent considérée comme un symbole de la lutte pour la préservation de la biodiversité ?",
                            "good_answer": "Le panda géant",
                            "answers": [
                                "Le tigre",
                                "Le dauphin",
                                "Le panda géant",
                                "Le koala"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande réserve naturelle du monde, située en Antarctique ?",
                            "good_answer": "La réserve naturelle de l''Antarctique",
                            "answers": [
                                "La réserve naturelle de l''Antarctique",
                                "La réserve naturelle du Gondwana",
                                "Le parc national de Yellowstone",
                                "Le parc national des Galápagos"
                            ]
                        },
                        {
                            "question": "Quel est le processus par lequel les espèces évoluent pour s''adapter à leur environnement ?",
                            "good_answer": "La sélection naturelle",
                            "answers": [
                                "La mutation génétique",
                                "La reproduction asexuée",
                                "L''ingénierie écologique",
                                "La sélection naturelle"
                            ]
                        }
                    ]
                },
                {
                    "id": 4,
                    "topic": "Biodiversite",
                    "title": "Les enjeux de la biodiversité",
                    "questions": [
                        {
                            "question": "Combien d''espèces animales sont répertoriées dans le monde ?",
                            "good_answer": "Environ 8,7 millions",
                            "answers": [
                                "Environ 1,8 million",
                                "Environ 8,7 millions",
                                "Environ 11,9 millions",
                                "Environ 15,2 millions"
                            ]
                        },
                        {
                            "question": "Quel pourcentage de la biodiversité marine est constitué de poissons ?",
                            "good_answer": "Environ 33%",
                            "answers": [
                                "Environ 25%",
                                "Environ 29%",
                                "Environ 33%",
                                "Environ 38%"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande menace pour la biodiversité actuellement ?",
                            "good_answer": "La perte d''habitat",
                            "answers": [
                                "La perte d''habitat",
                                "La chasse excessive",
                                "La pollution de l''air",
                                "La fragmentation des écosystèmes"
                            ]
                        },
                        {
                            "question": "Quel est l''objectif principal de la Convention sur la diversité biologique (CDB) ?",
                            "good_answer": "La conservation de la diversité biologique",
                            "answers": [
                                "La promotion de la biotechnologie",
                                "La gestion des déchets",
                                "La régulation des ressources génétiques",
                                "La conservation de la diversité biologique"
                            ]
                        },
                        {
                            "question": "Quel groupe d''animaux est le plus diversifié en termes d''espèces connues ?",
                            "good_answer": "Les insectes",
                            "answers": [
                                "Les mammifères",
                                "Les reptiles",
                                "Les poissons",
                                "Les insectes"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande zone de biodiversité terrestre au monde ?",
                            "good_answer": "La forêt amazonienne",
                            "answers": [
                                "Le désert du Sahara",
                                "La toundra arctique",
                                "La forêt amazonienne",
                                "Les prairies de Mongolie"
                            ]
                        },
                        {
                            "question": "Quel est le principal mécanisme de l''évolution de la biodiversité ?",
                            "good_answer": "La sélection naturelle",
                            "answers": [
                                "La sélection naturelle",
                                "La migration des espèces",
                                "La compétition interspécifique",
                                "Les mutations aléatoires"
                            ]
                        },
                        {
                            "question": "Quel pays abrite la plus grande diversité d''espèces d''oiseaux ?",
                            "good_answer": "La Colombie",
                            "answers": [
                                "Les États-Unis",
                                "La Colombie",
                                "L''Australie",
                                "La Chine"
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
                                "La proportion d''oxygène dans l''atmosphère terrestre",
                                "La variété des formes de vie sur Terre"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un écosystème ?",
                            "good_answer": "Un ensemble d''organismes vivants et de leur environnement physique",
                            "answers": [
                                "Un type de microscope",
                                "Une substance chimique toxique",
                                "Un ensemble d''organismes vivants et de leur environnement physique",
                                "Une unité de mesure de la pression atmosphérique"
                            ]
                        },
                        {
                            "question": "Quel est l''objectif de la conservation de la biodiversité ?",
                            "good_answer": "Préserver les espèces et les écosystèmes pour les générations futures",
                            "answers": [
                                "Préserver les espèces et les écosystèmes pour les générations futures",
                                "Détruire les habitats naturels",
                                "Encourager la perte irréversible d''espèces et d''écosystèmes",
                                "Promouvoir l''extinction des espèces"
                            ]
                        },
                        {
                            "question": "Quelle est la principale cause de la perte de biodiversité ?",
                            "good_answer": "La destruction des habitats naturels",
                            "answers": [
                                "Le réchauffement climatique",
                                "La migration des espèces",
                                "L''introduction d''espèces envahissantes",
                                "La destruction des habitats naturels"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''une espèce endémique ?",
                            "good_answer": "Une espèce présente uniquement dans une région spécifique",
                            "answers": [
                                "Une espèce présente uniquement dans une région spécifique",
                                "Une espèce en voie de disparition",
                                "Une espèce invasive",
                                "Une espèce migratrice"
                            ]
                        },
                        {
                            "question": "Quels sont les avantages de la biodiversité ?",
                            "good_answer": "La fourniture de ressources alimentaires et médicinales",
                            "answers": [
                                "La pollution de l''environnement",
                                "La destruction des écosystèmes",
                                "La fourniture de ressources alimentaires et médicinales",
                                "La régulation naturelle des ravageurs et des maladies"

                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un corridor biologique ?",
                            "good_answer": "Une zone qui relie les habitats naturels et permet aux espèces de se déplacer",
                            "answers": [
                                "Une méthode de culture biologique sans l''utilisation de pesticides",
                                "Une zone qui relie les habitats naturels et permet aux espèces de se déplacer",
                                "Un outil utilisé pour mesurer la température",
                                "Un terme désignant un groupe d''oies"
                            ]
                        },
                        {
                            "question": "Qu''est-ce que la surexploitation des ressources ?",
                            "good_answer": "L''utilisation excessive des ressources naturelles au point de mettre en danger leur survie",
                            "answers": [
                                "La multiplication des ressources naturelles",
                                "La limitation de l''accès aux ressources",
                                "La gestion durable des ressources naturelles",
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
            "description" : "Le changement climatique, causé par l''activité humaine, altère durablement les conditions météorologiques mondiales et impacte l''environnement.",
            "quizzes": [
                {
                    "id": 6,
                    "topic": "Changement climatique",
                    "title": "Qu''est-ce-que le changement climatique",
                    "questions": [
                        {
                            "question": "Quelle est la principale cause du changement climatique actuel ?",
                            "good_answer": "Les émissions de gaz à effet de serre",
                            "answers": [
                                "L''activité volcanique",
                                "Les éruptions solaires",
                                "Les émissions de gaz à effet de serre",
                                "La variation naturelle du climat"
                            ]
                        },
                        {
                            "question": "Quel gaz à effet de serre est le plus préoccupant pour le climat ?",
                            "good_answer": "Le dioxyde de carbone (CO2)",
                            "answers": [
                                "Le méthane (CH4)",
                                "Le protoxyde d''azote (N2O)",
                                "L''ozone troposphérique (O3)",
                                "Le dioxyde de carbone (CO2)"
                            ]
                        },
                        {
                            "question": "Quel accord international vise à lutter contre le changement climatique en réduisant les émissions de gaz à effet de serre ?",
                            "good_answer": "L''Accord de Paris",
                            "answers": [
                                "L''accord de Paris",
                                "Le traité de Kyoto",
                                "Le protocole de Montréal",
                                "La déclaration de Copenhague"
                            ]
                        },
                        {
                            "question": "Quel phénomène climatique extrême est amplifié par le changement climatique ?",
                            "good_answer": "Les vagues de chaleur",
                            "answers": [
                                "Les blizzards",
                                "Les tornades",
                                "Les tempêtes tropicales",
                                "Les vagues de chaleur"
                            ]
                        },
                        {
                            "question": "Quelle est la principale conséquence du réchauffement climatique sur les océans ?",
                            "good_answer": "L''acidification des océans",
                            "answers": [
                                "L''augmentation du niveau de la mer",
                                "L''acidification des océans",
                                "La diminution des tempêtes",
                                "La désertification des fonds marins"
                            ]
                        },
                        {
                            "question": "Quelle région du monde est la plus vulnérable au changement climatique en raison de sa dépendance à l''agriculture ?",
                            "good_answer": "L''Afrique subsaharienne",
                            "answers": [
                                "L''Europe occidentale",
                                "L''Amérique du Nord",
                                "L''Afrique subsaharienne",
                                "L''Asie du Sud-Est"
                            ]
                        },
                        {
                            "question": "Quel est l''effet du changement climatique sur les glaciers et les calottes glaciaires ?",
                            "good_answer": "La fonte rapide",
                            "answers": [
                                "La fonte rapide",
                                "L''augmentation de la glace",
                                "La stabilisation",
                                "La diminution graduelle"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des forêts dans la lutte contre le changement climatique ?",
                            "good_answer": "Elles absorbent le dioxyde de carbone (CO2)",
                            "answers": [
                                "Elles émettent du méthane (CH4)",
                                "Elles n''ont aucun effet sur le climat",
                                "Elles absorbent le dioxyde de carbone (CO2)",
                                "Elles produisent de l''ozone (O3)"
                            ]
                        }
                    ]
                },
                {
                    "id": 7,
                    "topic": "Changement climatique",
                    "title": "Comment s''effectue le changement climatique",
                    "questions": [
                        {
                            "question": "Quelle est la principale cause du changement climatique?",
                            "good_answer": "Les émissions de gaz à effet de serre",
                            "answers": [
                                "Les éruptions volcaniques",
                                "Les rayons cosmiques",
                                "L''activité solaire",
                                "Les émissions de gaz à effet de serre"
                            ]
                        },
                        {
                            "question": "Quel gaz est le plus responsable de l''effet de serre?",
                            "good_answer": "Le dioxyde de carbone (CO2)",
                            "answers": [
                                "L''oxygène (O2)",
                                "L''azote (N2)",
                                "Le dioxyde de carbone (CO2)",
                                "L''hélium (He)"
                            ]
                        },
                        {
                            "question": "Qu''est-ce que l''effet de serre?",
                            "good_answer": "Le phénomène par lequel certaines substances emprisonnent la chaleur dans l''atmosphère terrestre.",
                            "answers": [
                                "Le phénomène par lequel certaines substances emprisonnent la chaleur dans l''atmosphère terrestre.",
                                "Un vent fort",
                                "Les fortes pluies",
                                "Un courant océanique"
                            ]
                        },
                        {
                            "question": "Quelle est la conséquence du réchauffement climatique sur les glaciers?",
                            "good_answer": "La fonte des glaciers",
                            "answers": [
                                "L''expansion des glaciers",
                                "La croissance des glaciers",
                                "La fonte des glaciers",
                                "La couleur des glaciers devient rouge"
                            ]
                        },
                        {
                            "question": "Quel est le principal secteur émetteur de gaz à effet de serre?",
                            "good_answer": "Le secteur de l''énergie",
                            "answers": [
                                "L''agriculture",
                                "Le secteur de la santé",
                                "Le secteur du transport",
                                "Le secteur de l''énergie"
                            ]
                        },
                        {
                            "question": "Quel accord international vise à lutter contre le changement climatique?",
                            "good_answer": "L''Accord de Paris",
                            "answers": [
                                "L''accord de Rome",
                                "L''accord de Paris",
                                "L''accord de Londres",
                                "L''accord de Prague"
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
                                "Absorber le dioxyde de carbone de l''atmosphère",
                                "La multiplication des cyclones tropicaux"
                            ]
                        }
                    ]
                },
                {
                    "id": 8,
                    "topic": "Changement climatique",
                    "title": "Comprendre le réchauffement climatique",
                    "questions": [
                        {
                            "question": "Quel est l''impact du réchauffement climatique sur la fréquence des phénomènes météorologiques extrêmes ?",
                            "good_answer": "Ils deviennent plus fréquents",
                            "answers": [
                                "Ils deviennent moins fréquents",
                                "Ils restent inchangés",
                                "Ils deviennent plus fréquents",
                                "Ils diminuent en intensité"
                            ]
                        },
                        {
                            "question": "Quelle est la principale source d''émission de gaz à effet de serre liée aux activités humaines ?",
                            "good_answer": "La combustion des énergies fossiles",
                            "answers": [
                                "La combustion des énergies fossiles",
                                "L''agriculture",
                                "La déforestation",
                                "Les déchets industriels"
                            ]
                        },
                        {
                            "question": "Quel phénomène climatique est associé au réchauffement des océans ?",
                            "good_answer": "L''élévation du niveau de la mer",
                            "answers": [
                                "L''élévation du niveau de la mer",
                                "La diminution des tempêtes",
                                "La formation d''icebergs",
                                "L''acidification des océans"
                            ]
                        },
                        {
                            "question": "Quel est le principal contributeur au trou dans la couche d''ozone, un problème lié au changement climatique ?",
                            "good_answer": "Les gaz réfrigérants",
                            "answers": [
                                "Les émissions de CO2",
                                "Les poussières atmosphériques",
                                "Les gaz réfrigérants",
                                "Les déchets industriels"
                            ]
                        },
                        {
                            "question": "Quel est l''effet du réchauffement climatique sur la biodiversité ?",
                            "good_answer": "La perte de diversité biologique",
                            "answers": [
                                "L''augmentation des espèces",
                                "La migration des espèces",
                                "L''expansion des aires de répartition des espèces",
                                "La perte de diversité biologique"
                            ]
                        },
                        {
                            "question": "Quelle est la principale conséquence du réchauffement climatique sur les régions polaires ?",
                            "good_answer": "La fonte accélérée des glaciers",
                            "answers": [
                                "La fonte accélérée des glaciers",
                                "L''augmentation de la banquise",
                                "La stabilisation des températures",
                                "L''expansion des calottes glaciaires"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des océans dans l''absorption du surplus de chaleur dû au réchauffement climatique ?",
                            "good_answer": "Ils agissent comme un réservoir thermique",
                            "answers": [
                                "Ils amplifient le réchauffement",
                                "Ils sont insensibles à la chaleur",
                                "Ils agissent comme un réservoir thermique",
                                "Ils absorbent l''excès de chaleur du réchauffement climatique."
                            ]
                        },
                        {
                            "question": "Quel gaz est libéré par la fonte du permafrost, contribuant au réchauffement climatique ?",
                            "good_answer": "Le méthane",
                            "answers": [
                                "Le méthane",
                                "Le dioxyde de carbone",
                                "L''oxygène",
                                "L''azote"
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "id": 4,
            "name": "Ecologie marine",
            "description" : "L''écologie marine explore les liens entre organismes marins et environnement, essentiels pour la préservation des océans.",
            "quizzes": [
                {
                    "id": 9,
                    "topic": "Ecologie marine",
                    "title": "Qu''est-ce-que l''écologie marine",
                    "questions": [
                        {
                            "question": "Quelle est la principale menace pour les récifs coralliens ?",
                            "good_answer": "Le réchauffement climatique",
                            "answers": [
                                "La pollution plastique",
                                "La surpêche",
                                "Le réchauffement climatique",
                                "La diminution du taux de salinité de l''eau de mer"
                            ]
                        },
                        {
                            "question": "Quel est le plus grand animal de tous les temps ?",
                            "good_answer": "La baleine bleue",
                            "answers": [
                                "Le requin blanc",
                                "L''orque",
                                "Le dauphin",
                                "La baleine bleue"
                            ]
                        },
                        {
                            "question": "Comment s''appelle la zone en pleine mer où la vie est très abondante ?",
                            "good_answer": "La zone pélagique",
                            "answers": [
                                "La zone pélagique",
                                "La zone benthique",
                                "La zone abyssale",
                                "La zone littorale"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un herbier marin ?",
                            "good_answer": "Une prairie sous-marine d''herbes marines",
                            "answers": [
                                "Un corail géant",
                                "Une prairie sous-marine d''herbes marines",
                                "Un poisson de grande taille",
                                "Un écosystème aquatique constitué principalement de champignons marins."
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un écosystème côtier ?",
                            "good_answer": "Un ensemble de communautés vivantes entre la terre et la mer",
                            "answers": [
                                "Un écosystème situé au fond de l''océan",
                                "Un espace sans vie marine",
                                "Un ensemble de communautés vivantes entre la terre et la mer",
                                "Un regroupement de plantes marines uniquement."
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un tsunami ?",
                            "good_answer": "Une série de vagues provoquée par un séisme sous-marin",
                            "answers": [
                                "Un mouvement régulier des marées",
                                "Un tourbillon marin",
                                "Le prout d''une balaine bleue",
                                "Une série de vagues provoquée par un séisme sous-marin"
                            ]
                        },
                        {
                            "question": "Que sont les espèces invasives marines ?",
                            "good_answer": "Des espèces non indigènes qui envahissent un écosystème marin",
                            "answers": [
                                "Des espèces endémiques protégées",
                                "Des espèces non indigènes qui envahissent un écosystème marin",
                                "Des espèces rares et menacées",
                                "Des espèces indigènes adaptées aux écosystèmes marins."
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un récif artificiel ?",
                            "good_answer": "Une structure créée par l''homme pour favoriser la biodiversité marine",
                            "answers": [
                                "Une structure créée par l''homme pour favoriser la biodiversité marine",
                                "Un récif naturel formé par des coraux",
                                "Un récif sans vie marine",
                                "Une formation rocheuse sous-marine résultant de processus géologiques naturels."
                            ]
                        }
                    ]
                },
                {
                    "id": 10,
                    "topic": "Écologie marine",
                    "title": "Les enjeux de l''écologie marine",
                    "questions": [
                        {
                            "question": "Quel est le plus grand contributeur à la pollution plastique des océans ?",
                            "good_answer": "Les déchets plastiques à usage unique",
                            "answers": [
                                "Les emballages alimentaires en carton",
                                "Les déchets métalliques",
                                "Les déchets plastiques à usage unique",
                                "Les bouteilles d''eau en verre"
                            ]
                        },
                        {
                            "question": "Quel est le principal danger pour les récifs coralliens dans le contexte de l''écologie marine ?",
                            "good_answer": "Le blanchissement corallien",
                            "answers": [
                                "Le blanchissement corallien",
                                "L''acidification des océans",
                                "La surpêche",
                                "La migration des tortues marines"
                            ]
                        },
                        {
                            "question": "Quelle est la principale menace pour les tortues marines en matière d''écologie marine ?",
                            "good_answer": "Les filets de pêche abandonnés",
                            "answers": [
                                "Le changement climatique",
                                "La destruction de l''habitat",
                                "La pollution lumineuse des zones côtières",
                                "Les filets de pêche abandonnés"
                            ]
                        },
                        {
                            "question": "Quel est l''effet de la marée noire sur l''écosystème marin ?",
                            "good_answer": "La contamination des espèces marines",
                            "answers": [
                                "La stimulation de la croissance des algues",
                                "L''amélioration de la biodiversité",
                                "La promotion de la pêche durable",
                                "La contamination des espèces marines"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des mangroves dans l''équilibre de l''écosystème marin ?",
                            "good_answer": "Elles servent de nurserie pour de nombreuses espèces marines",
                            "answers": [
                                "Elles servent de nurserie pour de nombreuses espèces marines",
                                "Elles sont des lieux de reproduction pour les méduses",
                                "Elles n''ont aucun impact sur l''écologie marine",
                                "Elles sont principalement des habitats pour les requins"
                            ]
                        },
                        {
                            "question": "Quelle est la principale source de surpêche dans le monde, menaçant l''écologie marine ?",
                            "good_answer": "La pêche illégale, non déclarée et non réglementée (INN)",
                            "answers": [
                                "La pêche artisanale",
                                "La pêche industrielle responsable",
                                "La pêche illégale, non déclarée et non réglementée (INN)",
                                "La cueillette de coquillages excessive"
                            ]
                        },
                        {
                            "question": "Quel est l''impact de l''acidification des océans sur les coquillages et les mollusques ?",
                            "good_answer": "Elle affaiblit leurs coquilles",
                            "answers": [
                                "Elle renforce leurs coquilles",
                                "Elle affaiblit leurs coquilles",
                                "Elle n''affecte pas les coquillages",
                                "Elle les rend plus résistants aux prédateurs"
                            ]
                        },
                        {
                            "question": "Quelle est la conséquence de la destruction des herbiers marins sur l''écologie marine ?",
                            "good_answer": "La perte d''habitats essentiels pour de nombreuses espèces",
                            "answers": [
                                "L''amélioration de la qualité de l''eau",
                                "L''augmentation de la productivité marine",
                                "La perte d''habitats essentiels pour de nombreuses espèces",
                                "La diminution des températures océaniques"
                            ]
                        }
                    ]
                },
                {
                    "id": 11,
                    "topic": "Écologie marine",
                    "title": "Préservez l''écologie marine",
                    "questions": [
                        {
                            "question": "Quel est le rôle des herbiers marins dans l''écosystème marin ?",
                            "good_answer": "Ils servent de nurserie et d''habitat pour de nombreuses espèces",
                            "answers": [
                                "Ils servent de nurserie et d''habitat pour de nombreuses espèces",
                                "Ils produisent de l''oxygène pour les poissons",
                                "Ils n''ont aucun impact sur l''écosystème marin",
                                "Ils agissent comme des filtres naturels en purifiant l''eau de mer"
                            ]
                        },
                        {
                            "question": "Quelle est la principale menace pour les requins dans le contexte de l''écologie marine ?",
                            "good_answer": "La surpêche et le commerce des ailerons de requin",
                            "answers": [
                                "La perte d''habitats",
                                "Les changements climatiques",
                                "La surpêche et le commerce des ailerons de requin",
                                "La prédation par les méduses géantes"
                            ]
                        },
                        {
                            "question": "Quel est l''effet des marées rouges sur la faune marine ?",
                            "good_answer": "Elles peuvent causer des mortalités massives de poissons et de coquillages",
                            "answers": [
                                "Elles favorisent la croissance des algues",
                                "Elles n''ont aucun impact sur la faune marine",
                                "Elles stimulent la reproduction des espèces marines",
                                "Elles peuvent causer des mortalités massives de poissons et de coquillages"
                            ]
                        },
                        {
                            "question": "Quel est l''impact de la pollution sonore sur les cétacés, tels que les baleines ?",
                            "good_answer": "Elle perturbe leur communication et leur orientation",
                            "answers": [
                                "Elle renforce leur capacité d''écholocation",
                                "Elle perturbe leur communication et leur orientation",
                                "Elle n''affecte pas les cétacés",
                                "Elle améliore leur capacité à détecter les prédateurs marins"
                            ]
                        },
                        {
                            "question": "Quelle est la conséquence de la pollution par les hydrocarbures sur les oiseaux marins ?",
                            "good_answer": "Elle peut causer des dommages aux plumes et à la structure des plumes",
                            "answers": [
                                "Elle améliore la flottabilité des oiseaux",
                                "Elle n''a aucun effet sur les oiseaux marins",
                                "Elle stimule la croissance des plumes chez les oiseaux marins",
                                "Elle peut causer des dommages aux plumes et à la structure des plumes"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des récifs coralliens dans la protection des côtes contre les tempêtes ?",
                            "good_answer": "Ils agissent comme des barrières naturelles",
                            "answers": [
                                "Ils servent de sources d''énergie pour les tempêtes",
                                "Ils amplifient l''impact des tempêtes",
                                "Ils n''ont aucun effet sur les tempêtes",
                                "Ils agissent comme des barrières naturelles"
                            ]
                        },
                        {
                            "question": "Quel est l''effet de la montée du niveau de la mer sur les écosystèmes côtiers ?",
                            "good_answer": "Elle entraîne la perte d''habitats côtiers",
                            "answers": [
                                "Elle entraîne la perte d''habitats côtiers",
                                "Elle favorise la biodiversité côtière",
                                "Elle n''a aucun impact sur les écosystèmes côtiers",
                                "Elle stimule la croissance des écosystèmes côtiers"
                            ]
                        },
                        {
                            "question": "Quel est le principal problème lié à la pêche fantôme dans l''écologie marine ?",
                            "good_answer": "Les engins de pêche abandonnés qui continuent de capturer des animaux marins",
                            "answers": [
                                "La pêche excessive",
                                "Les engins de pêche abandonnés qui continuent de capturer des animaux marins",
                                "La destruction des habitats marins",
                                "La surpêche illégale"
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "id": 5,
            "name": "Protection des zones humides",
            "description" : "Protéger les zones humides préserve des écosystèmes essentiels, assurant eau douce, régulation des inondations et biodiversité unique.",
            "quizzes": [
                {
                    "id": 12,
                    "topic": "Protection des zones humides",
                    "title": "L''envol du Héron",
                    "questions": [
                        {
                            "question": "Quelle est la principale menace pour les zones humides ?",
                            "good_answer": "Le drainage excessif",
                            "answers": [
                                "Les castors",
                                "Les chauves-souris",
                                "Le drainage excessif",
                                "La surpopulation de grenouilles"
                            ]
                        },
                        {
                            "question": "Quel est le rôle des zones humides dans l''écosystème ?",
                            "good_answer": "Elles filtrent les eaux et purifient l''air",
                            "answers": [
                                "Elles sont des zones de recharge pour les extraterrestres",
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
                                "En introduisant des espèces de poissons dans l''écosystème des zones humides",
                                "En limitant les activités humaines impactantes"
                            ]
                        },
                        {
                            "question": "Quel oiseau est emblématique des zones humides ?",
                            "good_answer": "Le héron cendré",
                            "answers": [
                                "Le héron cendré",
                                "Le manchot empereur",
                                "Le pigeon voyageur",
                                "La mouette rieuse"
                            ]
                        },
                        {
                            "question": "Quel est l''effet des zones humides sur les inondations ?",
                            "good_answer": "Elles les atténuent en absorbant l''eau en excès",
                            "answers": [
                                "Elles les atténuent en absorbant l''eau en excès",
                                "Elles les aggravent en utilisant des pompiers",
                                "Elles les ignorent et font des châteaux de sable",
                                "Elles intensifient les inondations en organisant des fêtes aquatiques"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un marais ?",
                            "good_answer": "Un écosystème humide avec une végétation aquatique flottante",
                            "answers": [
                                "Une tartine beurrée",
                                "Une danse traditionnelle",
                                "Un écosystème humide avec une végétation aquatique flottante",
                                "Un nuage de poussière stellaire"
                            ]
                        },
                        {
                            "question": "Pourquoi est-il important de préserver les zones humides ?",
                            "good_answer": "Elles abritent une biodiversité exceptionnelle",
                            "answers": [
                                "Elles sont le repaire des extraterrestres",
                                "Elles ont des restaurants étoilés Michelin",
                                "Elles sont les quartiers généraux secrets des super-héros",
                                "Elles abritent une biodiversité exceptionnelle"
                            ]
                        },
                        {
                            "question": "Quelle est la meilleure façon de visiter une zone humide ?",
                            "good_answer": "En respectant la réglementation en vigueur",
                            "answers": [
                                "En organisant une rave-party géante",
                                "En respectant la réglementation en vigueur",
                                "En portant un sombrero fluorescent",
                                "En parcourant la zone en sautant d''une pierre à l''autre pour éviter l''eau"
                            ]
                        }
                    ]
                },
                {
                    "id": 13,
                    "topic": "Protection des zones humides",
                    "title": "Qu''est-ce-que la protection des zones humides ?",
                    "questions": [
                        {
                            "question": "Quelle est la principale cause de dégradation des zones humides ?",
                            "good_answer": "La pollution",
                            "answers": [
                                "Les oiseaux",
                                "La danse de la pluie",
                                "La pollution",
                                "Le réchauffement climatique"
                            ]
                        },
                        {
                            "question": "Quel est l''animal emblématique des zones humides ?",
                            "good_answer": "Le castor",
                            "answers": [
                                "La licorne",
                                "Le castor",
                                "Le kangourou",
                                "Le tigre"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''un marais ?",
                            "good_answer": "Une zone humide principalement composée d''herbes et de roseaux",
                            "answers": [
                                "Une zone humide principalement composée d''herbes et de roseaux",
                                "Un dessert sucré",
                                "Un type de chapeau",
                                "un quartier dans Paris"
                            ]
                        },
                        {
                            "question": "Quel est l''effet bénéfique des zones humides sur l''environnement ?",
                            "good_answer": "Elles filtrent et purifient l''eau",
                            "answers": [
                                "Elles rendent les grenouilles sympathiques",
                                "Elles produisent des arcs-en-ciel",
                                "Elles filtrent et purifient l''eau",
                                "Elles servent de piste de danse pour les oiseaux migrateurs"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande zone humide du monde ?",
                            "good_answer": "Le Pantanal au Brésil",
                            "answers": [
                                "Mon jardin",
                                "Le désert du Sahara",
                                "La Seine",
                                "Le Pantanal au Brésil"
                            ]
                        },
                        {
                            "question": "Comment appelle-t-on un petit étang formé dans une zone humide ?",
                            "good_answer": "Une mare",
                            "answers": [
                                "Un bouillon de culture",
                                "Un jacuzzi",
                                "Une mare",
                                "Un bain de grenouilles"
                            ]
                        },
                        {
                            "question": "Quelle est la couleur de la vase présente dans les zones humides ?",
                            "good_answer": "Marron",
                            "answers": [
                                "Marron",
                                "Orange fluo",
                                "Arc-en-ciel",
                                "Vert"
                            ]
                        },
                        {
                            "question": "Qu''est-ce qu''une tourbière ?",
                            "good_answer": "Une zone humide où se forme progressivement une sorte de terre spéciale appelée tourbe",
                            "answers": [
                                "Une tour en béton",
                                "Une espèce de biscuit",
                                "Un type de fruit tropical à la chair juteuse et sucrée",
                                "Une zone humide où se forme progressivement une sorte de terre spéciale appelée tourbe"
                            ]
                        }
                    ]
                },
                {
                    "id": 14,
                    "topic": "Protection des zones humides",
                    "title": "Le marécage sophistiqué",
                    "questions": [
                        {
                            "question": "Quel pourcentage des zones humides mondiales ont disparu au cours des 100 dernières années ?",
                            "good_answer": "50%",
                            "answers": [
                                "5%",
                                "25%",
                                "50%",
                                "75%"
                            ]
                        },
                        {
                            "question": "Quelle est la plus grande zone humide du monde en termes de superficie ?",
                            "good_answer": "Le Pantanal",
                            "answers": [
                                "Le Pantanal",
                                "Le Marais Poitevin",
                                "Les Everglades",
                                "La Savane Mara"
                            ]
                        },
                        {
                            "question": "Combien d''espèces d''oiseaux dépendent des zones humides pour leur survie ?",
                            "good_answer": "Plus de 1 000",
                            "answers": [
                                "Moins de 100",
                                "Environ 500",
                                "Plus de 1 000",
                                "Plus de 10 000"
                            ]
                        },
                        {
                            "question": "Quelle proportion des espèces menacées d''extinction vivent dans les zones humides ?",
                            "good_answer": "Environ 40%",
                            "answers": [
                                "Moins de 10%",
                                "Environ 20%",
                                "Environ 30%",
                                "Environ 40%"
                            ]
                        },
                        {
                            "question": "Quelle est la principale cause de la destruction des zones humides dans le monde ?",
                            "good_answer": "L''urbanisation",
                            "answers": [
                                "Le réchauffement climatique",
                                "L''urbanisation",
                                "L''agriculture intensive",
                                "La surpêche"
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "id": 6,
            "name": "Conservation de l''eau",
            "description" : "La conservation de l''eau vise à économiser et gérer judicieusement cette ressource pour répondre aux besoins présents et futurs.",
            "quizzes": [
                {
                    "id": 15,
                    "topic": "Conservation de l''eau",
                    "title": "Le joueur d''eau",
                    "questions": [
                            {
                                "question": "Quelle est la quantité d''eau nécessaire pour qu''un humain survive pendant une semaine sans aucune autre source de nourriture ?",
                                "good_answer": "7 litres",
                                "answers": [
                                    "1 litre",
                                    "7 litres",
                                    "13 litres",
                                    "20 litres"
                                ]
                            },
                            {
                                "question": "Quelle est la principale cause de perte d''eau potable dans les fuites d''un système de distribution d''eau ?",
                                "good_answer": "Les canalisations défectueuses",
                                "answers": [
                                    "Les robinets qui restent ouverts",
                                    "L''évaporation naturelle",
                                    "Les canalisations défectueuses",
                                    "Les poissons qui boivent l''eau"
                                ]
                            },
                            {
                                "question": "Quelle partie du monde est le plus touchée par le stress hydrique ?",
                                "good_answer": "Le Moyen-Orient",
                                "answers": [
                                    "L''Amérique du Nord",
                                    "L''Australie",
                                    "L''Antarctique",
                                    "Le Moyen-Orient"
                                ]
                            },
                            {
                                "question": "Quelle est la proportion d''eau salée sur Terre ?",
                                "good_answer": "97 %",
                                "answers": [
                                    "50 %",
                                    "78 %",
                                    "91 %",
                                    "97 %"
                                ]
                            },
                            {
                                "question": "Dans quelle activité quotidienne consommons-nous le plus d''eau en moyenne ?",
                                "good_answer": "L''agriculture",
                                "answers": [
                                    "L''agriculture",
                                    "La consommation domestique",
                                    "L''industrie",
                                    "La pêche sportive"
                                ]
                            }
                        ]
                },
                {
                    "id": 16,
                    "topic": "Conservation de l''eau",
                    "title": "Plouf plouf à la rescousse",
                    "questions": [
                            {
                                "question": "Quelle est la quantité d''eau douce disponible sur Terre ?",
                                "good_answer": "2,5%",
                                "answers": [
                                    "2,5 %",
                                    "3,4 %",
                                    "5,3 %",
                                    "5,7 %"
                                ]
                            },
                            {
                                "question": "Quelle est la consommation moyenne d''eau d''une chasse d''eau traditionnelle ?",
                                "good_answer": "9 à 12 litres",
                                "answers": [
                                    "2 à 4 litres",
                                    "9 à 12 litres",
                                    "15 à 18 litres",
                                    "21 à 25 litres"
                                ]
                            },
                            {
                                "question": "Quelle activité humaine utilise le plus d''eau dans le monde ?",
                                "good_answer": "L''agriculture",
                                "answers": [
                                    "La production d''électricité",
                                    "La consommation domestique",
                                    "L''agriculture",
                                    "Les parcs aquatiques"
                                ]
                            },
                            {
                                "question": "Combien de litres d''eau faut-il pour produire une seule canette de soda ?",
                                "good_answer": "250 à 300 litres",
                                "answers": [
                                    "100 à 150 litres",
                                    "150 à 200 litres",
                                    "200 à 250 litres",
                                    "250 à 300 litres"
                                ]
                            },
                            {
                                "question": "Quel est le pourcentage d''eau froide dans un iceberg ?",
                                "good_answer": "90%",
                                "answers": [
                                    "85 %",
                                    "88 %",
                                    "90 %",
                                    "97 %"
                                ]
                            },
                            {
                                "question": "Quel pays possède le plus grand pourcentage de sa population sans accès à l''eau potable ?",
                                "good_answer": "Somalie",
                                "answers": [
                                    "Inde",
                                    "Brésil",
                                    "Somalie",
                                    "Botswana"
                                ]
                            },
                            {
                                "question": "Combien de litres d''eau faut-il pour produire une tonne de papier recyclé ?",
                                "good_answer": "250 000 litres",
                                "answers": [
                                    "100 000 litres",
                                    "150 000 litres",
                                    "200 000 litres",
                                    "250 000 litres"
                                ]
                            }
                        ]
                },
                {
                    "id": 17,
                    "topic": "Conservation de l''eau",
                    "title": "Pourquoi l''eau ne dit jamais de blagues ?",
                    "questions": [
                            {
                                "question": "Quelle est la principale source d''eau douce utilisée par les humains dans le monde ?",
                                "good_answer": "Les eaux souterraines",
                                "answers": [
                                    "Les rivières",
                                    "Les lacs",
                                    "Les eaux souterraines",
                                    "Les glaciers"
                                ]
                            },
                            {
                                "question": "Combien de litres d''eau utilise-t-on en moyenne par jour pour une douche de 10 minutes ?",
                                "good_answer": "Environ 80 litres",
                                "answers": [
                                    "Environ 20 litres",
                                    "Environ 50 litres",
                                    "Environ 80 litres",
                                    "Environ 120 litres"
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
                },
                {
                    "id": 18,
                    "topic": "Conservation de l''eau",
                    "title": "Quand une goutte d''eau se transforme en super-héros",
                    "questions": [
                            {
                                "question": "Quelle est la principale cause de la pénurie d''eau dans de nombreuses régions du monde ?",
                                "good_answer": "Le gaspillage excessif d''eau",
                                "answers": [
                                    "La pollution de l''eau",
                                    "Le gaspillage excessif d''eau",
                                    "Les précipitations insuffisantes",
                                    "La surconsommation de boissons gazeuses"
                                ]
                            },
                            {
                                "question": "Quel pourcentage de l''eau sur Terre est disponible pour la consommation humaine ?",
                                "good_answer": "Moins de 1%",
                                "answers": [
                                    "Moins de 1%",
                                    "Environ 25%",
                                    "Près de 50%",
                                    "Plus de 75%"
                                ]
                            },
                            {
                                "question": "Qu''est-ce que l''agriculture durable peut contribuer à la conservation de l''eau ?",
                                "good_answer": "La mise en place de techniques d''irrigation efficaces",
                                "answers": [
                                    "L''abandon de l''agriculture",
                                    "L''utilisation intensive d''engrais chimiques",
                                    "La mise en place de techniques d''irrigation efficaces",
                                    "L''augmentation de la consommation d''eau dans les exploitations agricoles"
                                ]
                            }
                        ]
                }
            ]
        }
    ],
    "version": 1
}', NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM quiz_games
WHERE gid = (SELECT id FROM games WHERE name = 'Quiz' AND game_version = '1.0');
-- +goose StatementEnd
