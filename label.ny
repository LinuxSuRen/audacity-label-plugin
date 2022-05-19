$nyquist plug-in
$version 4
$type tool analyze
$name "Label export"

$control filename (_ "Export Label Track to File:") file (_ "Select a file") "*default*/subtitles.md" "Markdown file|*.md" "save,overwrite"


;;; Reurn number as string with at least 2 digits
(defun pad (num)
  (format nil "~a~a" (if (< num 10) "0" "") num))


;;; Format time (seconds) as mm:ss
(defun md-time-format (sec)
  (let* ((seconds (truncate sec))
         (mm (truncate (/ seconds 60)))
         (ss (rem seconds 60)))
    (format nil "~a:~a" (pad mm) (pad ss))))


;;; generate mp3 lyric
;;;;;; srt sample text:
;;; [00:00.26] subtitle 1
;;; [00:02.22] subtitle 2
(defun label-to-md (labels)
  (setf md "")

  (dolist (label labels)
    (setf timeS (md-time-format (first label)))
    (string-append md (format nil "[~a] ~a~%" timeS (third label))))
  (format nil md))


;;; Return file extension or empty string
(defun get-file-extension (fname)
  (let ((n (1- (length fname)))
        (ext ""))
    (do ((i n (1- i)))
        ((= i 0) ext)
      (when (char= (char fname i) #\.)
        (setf ext (subseq fname i))
        (return ext)))))


(setf file-ext (string-upcase (get-file-extension filename)))

;; Get labels from first label track
(setf labels (second (first (aud-get-info "labels"))))

;; detect file extension to determine which format to export
(setf txt (if (string= ".md" file-ext)
            (label-to-md labels)
            (label-to-md labels)))

(setf fp (open filename :direction :output))
(format fp txt)
(close fp)
