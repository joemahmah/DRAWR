
# DRAWR (Drawing Really Awful Works of aRt)
Meet DRAWR, the really bad artist. DRAWR learns to draw by looking at images. Sadly, while he may have a good memory, he doesn't really understand what actually makes an image art... DRAWR records the colors of neighboring pixels and uses a Markov algorithm to generate new images.

Some things that DRAWR can and can't do:
* DRAWR is able to generate images.
* DRAWR is able to learn how to draw from images (jpg, png).
* DRAWR is NOT able to actually understand what is sees in images (i.e. it cannot understand what a head is).

## The Future of DRAWR
I'd really like to expand the scope of DRAWR. Currently, the program uses a simple Markov algorithm to generat images.

First, I'd like to expand the depth of the markov chain. At a depth of 1, the algorithm only looks at the pixel to the left and above when generating a new point. At a depth of n, the algorithm would look left and up in lines of length n pixels.

Second, I want to add an optional web interface for DRAWR. This will provide an easy to use GUI that will also be able to actually render the images generated. Not to mention, it will make it easier to host on a remote server (Note: DRAWR's web interface will likely not be very secure or complex. It is not really intended to be used as a functioning service. If you desire to have DRAWR publically accessible, it will be reccomended to have another service handle routing traffic to DRAWR).

Third, I would like to add an option to have DRAWR generate images based on square chunks rather than lines. This would be really slow, but it would likely make DRAWR generate much better looking images.

Next, I'd like to add the ability for DRAWR to pull images a url. First itterations would likely require the url to link directly to the image, but later versions would probably be able to search a page for images.

Finally, I'd like to add a mode to let DRAWR look at source images more analytically. Namely, DRAWR would find groups of similarly colored pixels and store the geometry along with color data. The geometries would be categorized as "similar looking" based on size and shape. Perhaps, DRAWR would ask for the user to define tags for the geometries? DRAWR would then decide on a starting geometry and use said geometry to generate more geometries in the image. Once DRAWR has finished with geometries, it would fill the geometries and decide upon colorization for the rest of the image using the combined data of the geometries. This would hopefully allow for images that actually look like something...

## Compiling
DRAWR is written in Go (so you need a go compiler). DRAWR does not have any external dependencies.

A simple `go build boot.go` should spit out an executable.

## License
DRAWR is licensed under an MIT license.
