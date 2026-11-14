#!/usr/bin/env python3
"""Renders the Gedanken icon set: hydrogen's 2p orbital as a glowing cloud
of electron positions sampled from |psi|^2, blue and amber for the sign
of psi, on the app's night-blue backdrop.

Run: python3 tools/make_icon.py assets   (needs numpy and Pillow)

Outputs (1024 px): icon.png (iOS and legacy Android), android-foreground.png,
android-background.png, android-monochrome.png."""
import sys
import numpy as np
from PIL import Image

out = sys.argv[1]
N = 2048
rng = np.random.default_rng(11)
NIGHT = np.array([10, 13, 22]) / 255          # #0a0d16, the app and splash colour
GLOW = np.array([0.10, 0.14, 0.30])

def blur(img, sigma):
    """Gaussian blur by FFT."""
    f = np.fft.fftfreq(N)
    g = np.exp(-2 * (np.pi * sigma) ** 2 * (f[:, None] ** 2 + f[None, :] ** 2))
    return np.real(np.fft.ifft2(np.fft.fft2(img) * g))

def sample_2p(n, extent=12.0):
    peak = (2 * np.exp(-1)) ** 2
    out = []
    while sum(len(p) for p in out) < n:
        xyz = rng.uniform(-extent, extent, size=(200000, 3))
        r = np.linalg.norm(xyz, axis=1)
        psi = xyz[:, 2] * np.exp(-r / 2)
        keep = (r < extent) & (rng.uniform(0, peak, len(r)) < psi ** 2)
        out.append(xyz[keep])
    return np.concatenate(out)[:n]

pts = sample_2p(3200)
tilt = np.radians(-22)                        # lean the lobe axis for some life
x, z = pts[:, 0], pts[:, 2]
px0 = x * np.cos(tilt) - z * np.sin(tilt)
pz0 = x * np.sin(tilt) + z * np.cos(tilt)
depth = pts[:, 1]
positive = pts[:, 2] >= 0
reach = np.percentile(np.hypot(px0, pz0), 96)
yy, xx = np.mgrid[0:N, 0:N]
dist = np.hypot(xx - N / 2, yy - N / 2) / (N / 2)

def light(frac):
    """Emitted light (RGB, unbounded) and coverage for the cloud filling frac of the canvas."""
    s = N * frac / 2 / reach
    cols = {True: np.array([0.55, 0.73, 1.0]), False: np.array([1.0, 0.68, 0.42])}
    rgb = np.zeros((N, N, 3))
    cover = np.zeros((N, N))
    for sign, col in cols.items():
        m = positive == sign
        w = 0.55 + 0.45 * (depth[m] - depth.min()) / np.ptp(depth)  # nearer dots brighter
        h, _, _ = np.histogram2d(N / 2 - pz0[m] * s, N / 2 + px0[m] * s, bins=N,
                                 range=[[0, N], [0, N]], weights=w)
        dots = blur(h, 2.6 * frac / 0.78) * 95 * (frac / 0.78) ** 2
        haze = blur(h, 60 * frac / 0.78) * 480 * (frac / 0.78) ** 2
        e = np.clip(dots, 0, None) + np.clip(haze, 0, None)
        rgb += e[..., None] * col
        cover += e
    core = 1.4 * np.exp(-(dist / (0.014 * frac / 0.78)) ** 2) + 0.25 * np.exp(-(dist / (0.06 * frac / 0.78)) ** 2)
    rgb += core[..., None] * np.array([1.0, 0.97, 0.9])
    cover += core
    return rgb, cover

def tone(rgb):
    return 1 - np.exp(-rgb)

def background():
    return NIGHT + (GLOW - NIGHT) * np.clip(1 - dist, 0, 1)[..., None] ** 2.2

def save(arr, mode, name):
    Image.fromarray((np.clip(arr, 0, 1) * 255).astype(np.uint8), mode).resize((1024, 1024), Image.LANCZOS).save(f"{out}/{name}")

# Full icon: cloud on the backdrop.
rgb, _ = light(0.74)
bg = background()
save(1 - (1 - bg) * (1 - tone(rgb)), "RGB", "icon.png")

# Android adaptive layers: artwork inside the middle 66 of 108 dp.
frac = 0.74 * 66 / 108
rgb, cover = light(frac)
alpha = tone(cover)
color = np.clip(tone(rgb) / np.maximum(alpha[..., None], 1e-6), 0, 1)
save(np.dstack([color, alpha]), "RGBA", "android-foreground.png")
save(background(), "RGB", "android-background.png")
save(np.dstack([np.ones((N, N, 3)), alpha]), "RGBA", "android-monochrome.png")
print("done")
