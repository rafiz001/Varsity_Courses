# Mid Term
## Lab-1: Color Space Conversion
Write a Python program to perform the following tasks:
- a) Read an image.
- b) Display the image resolution in the format: Resolution = Width × Height pixels.
- c) Print the image color mode.
- d) Convert the image into grayscale.
- e) Display both the original image and the grayscale image.
- f) Label each displayed image with an appropriate title.
- g) Save the grayscale image.
## Lab-2: Brightness adjustment of grayscale images
Write a Python program to perform the following tasks:
- a) Read a grayscale image.
- b) Increase the brightness of the image by multiplying each pixel value by a scaling factor S₁ = 1.5
- c) Decrease the brightness of the image by multiplying each pixel value by a scaling factor S₂ = 0.3
- d) Compare the original grayscale image with the brightness-adjusted images.

## Lab-3: Image Thresholding
Write a Python program to perform the following tasks:
- a) Read a grayscale image
- b) Set a threshold value, k = 128
- c) Apply binary thresholding using:
if r > k → s = 255
else → s = 0
- d) Display the binary image
- e) Create three binary images using different thresholds: k = 50, 100, 150
- f) Compare the results
## Lab-4: Image Negative Transformation
Write a Python program to perform the following tasks:
- a) Read a grayscale image
- b) Apply image negative transformation using the formula:  s = 255 – r
- c) Display both images

# Lab-5: Log Transformation

Write a Python program to perform the following tasks:
- a) Read a grayscale image.
- b) Apply log transformation using the formula:  
\[ s = c \times \log_{10}(1 + r) \]
- c) Display both the original image and the log-transformed image with appropriate titles.

# Lab-6: Power-Law Transformation

Write a Python program to perform the following tasks:
- a) Read a grayscale image.
- b) Apply power-law transformation using the formula:  
\[ s = c \times r^\gamma \]
- c) Use at least two different values of \( \gamma \) and compare the results.
- d) Display the original image and the transformed images with appropriate titles.

# Lab-7: Gray Level Slicing

Write a Python program to perform the following tasks:
- a) Read a grayscale image.
- b) Define a gray-level range [A, B], where A=100 and B=180.
- c) Perform gray-level slicing with background preserved using the following condition:\
   If \( A \leq r \leq B \), then set \( s = 255 \)
   \
   Otherwise, keep the original pixel value
- d) Perform gray-level slicing without background (background suppressed) using the following condition:\
   If \( A \leq r \leq B \), then set \( s = 220 \)
   \
   Otherwise, set \( s = 0 \)
- e) Display all three images (Original image and 2 gray level sliced image) with appropriate titles.