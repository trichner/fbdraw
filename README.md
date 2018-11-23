
Those are the chronicles of my search for knowledge about Linux framebuffer devices.
You can find the devices in `/dev/fb{0..32}`.

## Double-Buffering
The proper term is 'panning'. The `FBIOGET_FSCREENINFO` `ioctl` indicates if this is supported by the driver or not.

https://stackoverflow.com/questions/13907471/linux-framebuffer-graphics-and-vsync


## 4DPi-24-HAT Properties
2.4" HAT Primary Display for the Raspberry Pi

```
{ // FBIOGET_FSCREENINFO
  Id:[108 99 100 112 105 0 0 0 0 0 0 0 0 0 0 0]
  Smem_start:3154890752
  Smem_len:155648
  Type:0
  Type_aux:0
  Visual:2
  Xpanstep:0
  Ypanstep:0
  Ywrapstep:0
  Line_length:640
  Mmio_start:0
  Mmio_len:0
  Accel:0
  Capabilities:0
  Reserved:[0 0]
}
{ // FBIOGET_VSCREENINFO
  Xres:320
  Yres:240
  Xres_virtual:320
  Yres_virtual:240
  Xoffset:0
  Yoffset:0
  Bits_per_pixel:16
  Grayscale:0
  Red:{Offset:11 Length:5 Msb_right:0}
  Green:{Offset:5 Length:6 Msb_right:0}
  Blue:{Offset:0 Length:5 Msb_right:0}
  Transp:{Offset:0 Length:0 Msb_right:0}
  Nonstd:0
  Activate:0
  Height:49
  Width:74
  Accel_flags:0
  Pixclock:0
  Left_margin:0
  Right_margin:0
  Upper_margin:0
  Lower_margin:0
  Hsync_len:0
  Vsync_len:0
  Sync:0
  Vmode:0
  Rotate:0
  Colorspace:0
  Reserved:[0 0 0 0]
}
```