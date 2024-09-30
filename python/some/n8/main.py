import cv2

cam_id = 6489
f = open("urls.txt", "w")

# ------------------------------------------------------------------------

count = 0
while cam_id < 10_000:
    id_str = str(cam_id).zfill(4)
    url = f"https://krkvideo1.orionnet.online/cam{id_str}/tracks-v1/mono.m3u8"
    print(f"url: {url}")

    cap = cv2.VideoCapture(url)
    if cap.isOpened():
        print(f"opened: {id_str}")
        f.write(f"{count}. {url}\n")
        f.flush()
        
        count += 1
        cap.release()
    else:
        print(f"skipped: {id_str}")

    cam_id += 1

f.close()
exit(0)

# ------------------------------------------------------------------------

while True:
    ret, imag = cap.read()
    if ret:
        imag_out=imag.copy()
        gray = cv2.cvtColor(imag, cv2.COLOR_BGR2GRAY)
        cv2.imshow("Camera",imag)
        if cv2.waitKey(10) & 0xFF == 27:
            break
cv2.destroyAllWindows()
cap.release()