# Mid Term

---

## Lab-1: Color Space Conversion

Write a Python program to perform the following tasks:

- a) Read an image.  
- b) Display the image resolution in the format:  
  **Resolution = Width × Height pixels**  
- c) Print the image color mode.  
- d) Convert the image into grayscale.  
- e) Display both the original image and the grayscale image.  
- f) Label each displayed image with an appropriate title.  
- g) Save the grayscale image.  

---

## Lab-2: Brightness Adjustment of Grayscale Images

Write a Python program to perform the following tasks:

- a) Read a grayscale image.  
- b) Increase the brightness of the image by multiplying each pixel value by a scaling factor `S₁ = 1.5`.  
- c) Decrease the brightness of the image by multiplying each pixel value by a scaling factor `S₂ = 0.3`.  
- d) Compare the original grayscale image with the brightness-adjusted images.  

---

## Lab-3: Image Thresholding

Write a Python program to perform the following tasks:

- a) Read a grayscale image.  
- b) Set a threshold value `k = 128`.  
- c) Apply binary thresholding using:

```python
if r > k:
    s = 255
else:
    s = 0
```

- d) Display the binary image.  
- e) Create three binary images using different thresholds:  
  - `k = 50`  
  - `k = 100`  
  - `k = 150`  
- f) Compare the results.  

---

## Lab-4: Image Negative Transformation

Write a Python program to perform the following tasks:

- a) Read a grayscale image.  
- b) Apply image negative transformation using the formula:

```python
s = 255 - r
```

- c) Display both images.  

---

## Lab-5: Log Transformation

Write a Python program to perform the following tasks:

- a) Read a grayscale image.  
- b) Apply log transformation using the formula:

```python
s = c * log10(1 + r)
```

- c) Display both the original image and the log-transformed image with appropriate titles.  

---

## Lab-6: Power-Law Transformation

Write a Python program to perform the following tasks:

- a) Read a grayscale image.  
- b) Apply power-law transformation using the formula:

```python
s = c * r^gamma
```

- c) Use at least two different values of `gamma` and compare the results.  
- d) Display the original image and the transformed images with appropriate titles.  

---

## Lab-7: Gray Level Slicing

Write a Python program to perform the following tasks:

- a) Read a grayscale image.  
- b) Define a gray-level range `[A, B]`, where `A = 100` and `B = 180`.  

- c) Perform gray-level slicing **with background preserved** using the following condition:

```python
if A <= r <= B:
    s = 255
else:
    s = r
```

- d) Perform gray-level slicing **without background (background suppressed)** using the following condition:

```python
if A <= r <= B:
    s = 220
else:
    s = 0
```

- e) Display all three images (Original image and 2 gray-level sliced images) with appropriate titles.
---

## Lab-8: Bit Plane Slicing and Image Reconstruction
Write a Python program to perform the following tasks:
- a) Read a grayscale image.
- b) Apply bit plane slicing to extract all 8 bit planes (0–7) using the expression:
```
"bit_plane"_k = (img ">>" k) & 1
[render this math from https://asciimath.org/]
```
- c) Store all the bit planes into an array.
- d) Display all extracted bit planes in a 2 × 4 grid with appropriate titles (Bit Plane 0 to Bit Plane 7).
- e) Reconstruct a new image using only the higher-order bit planes (bit planes 4–7) using the formula:
```
s = sum_(k=4)^7 "bit_plane"_k*2^k
[render this math from https://asciimath.org/]
```
- f) Display both the original image and the reconstructed image with appropriate titles.
---

## Lab-9: Histogram Equalization
Write a Python program to perform histogram equalization on a grayscale image using the following steps:
- a) Compute frequency of each intensity level 
- b) Calculate probability density function (PDF) 
- c) Compute cumulative distribution function (CDF) 
- d) Map old pixel values to new values

## Lab-10: Image Smoothing

Write a Python program to perform image smoothing using spatial filters:

- a) Read a grayscale image
- b) Use replicate padding (edge padding) to handle boundary pixels
- c) Apply a 3×3 averaging filter to smooth the image
- d) Apply a 3×3 median filter to smooth the image
- e) Display the Original image, Averaging filter output, and Median filter output