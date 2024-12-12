import csv
from collections import Counter
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.preprocessing import LabelEncoder
import numpy as np

# Archivo de resultados
filename = "ruleta_resultados.csv"

# Función para cargar los datos del archivo CSV
def cargar_datos():
    try:
        with open(filename, mode="r") as file:
            reader = csv.reader(file)
            resultados = [row[0] for row in reader]
        return resultados
    except FileNotFoundError:
        print("No hay resultados registrados. Registra algunos antes de usar el modelo.")
        return []

# Función para entrenar el modelo
def entrenar_modelo(resultados):
    if len(resultados) < 10:  # Se necesitan al menos 10 datos para entrenar
        print("Se necesitan más datos para entrenar el modelo.")
        return None, None
    
    # Convertir los resultados en números (codificación)
    le = LabelEncoder()
    datos_codificados = le.fit_transform(resultados)

    # Crear características (X) y etiquetas (y)
    X = []
    for i in range(1, len(datos_codificados)):
        X.append([datos_codificados[i-1]])  # Usar el resultado anterior como característica
    y = datos_codificados[1:]  # El resultado actual es la etiqueta

    # Dividir los datos en entrenamiento y prueba
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # Entrenar el modelo
    modelo = RandomForestClassifier(random_state=42)
    modelo.fit(X_train, y_train)

    return modelo, le

# Función para predecir el próximo resultado
def predecir_resultado(modelo, le, ultimo_resultado):
    if modelo is None or le is None:
        print("El modelo no está entrenado.")
        return
    # Convertir el último resultado a un formato codificado
    ultimo_codificado = le.transform([ultimo_resultado])
    prediccion_codificada = modelo.predict([ultimo_codificado])
    prediccion = le.inverse_transform(prediccion_codificada)
    print(f"El modelo predice que el próximo resultado será: {prediccion[0]}")

# Función para ingresar un nuevo resultado
def ingresar_resultado():
    posicion = int(input("Introduce la posición de la aguja (en grados): "))
    resultado = input("Introduce el resultado (ejemplo: x0, x10, x25, etc.): ")
    with open(filename, mode="a", newline="") as file:
        writer = csv.writer(file)
        writer.writerow([resultado])
    print(f"Resultado '{resultado}' en la posición {posicion} registrado.")

# Menú principal
def menu():
    modelo, le = None, None
    while True:
        print("\n=== Ruleta ===")
        print("1. Ingresar nuevo resultado")
        print("2. Predecir próximo resultado")
        print("3. Salir")
        opcion = input("Selecciona una opción: ")

        if opcion == "1":
            ingresar_resultado()
            # Cargar y entrenar modelo después de ingresar nuevos resultados
            resultados = cargar_datos()
            modelo, le = entrenar_modelo(resultados)
        elif opcion == "2":
            if modelo is None or le is None:
                print("El modelo no está entrenado. Ingresa al menos un resultado primero.")
            else:
                ultimo_resultado = input("Introduce el último resultado (ejemplo: x0, x10, x25): ")
                predecir_resultado(modelo, le, ultimo_resultado)
        elif opcion == "3":
            print("¡Hasta luego!")
            break
        else:
            print("Opción inválida. Intenta de nuevo.")

# Ejecutar el menú
if __name__ == "__main__":
    menu()
