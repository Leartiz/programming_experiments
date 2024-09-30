import cv2

cap = cv2.VideoCapture("https://krkvideo1.orionnet.online/cam6500/tracks-v1/mono.m3u8")
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