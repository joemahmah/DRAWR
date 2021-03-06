
# DRAWR (Drawing Really Awful Works of aRt)
Meet DRAWR, the really bad artist. DRAWR learns to draw by looking at images. Sadly, while he may have a good memory, he doesn't really understand what actually makes an image art... DRAWR records the colors of neighboring pixels and uses a Markov algorithm to generate new images.

Some things that DRAWR can and can't do:
* DRAWR is able to generate images.
* DRAWR is able to learn how to draw from images (jpg, png).
* DRAWR is NOT able to actually understand what is sees in images (i.e. it cannot understand what a head is).

As a warning, DRAWR IS VERY SLOW. One of the issues with DRAWR is that there are about 4 billion pixel combinations (about 16.5 million combinations without alpha); add on the generation method and it becomes really slow. DRAWR will take a long time if you are generating a very large image or the source image has a large number of unique pixels (colors).

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

## Using DRAWR
DRAWR has three major components. Put together, they form the DRAWR executable. For the most part, you'll only need to really care for the generator compenent.

If desired, you can use these components without the DRAWR front end. 

### Generators
Generators are the compenents that actually generate the image. Generators are designed to support multiple generation modes. Different generators are able to utilize different types of containers. This is done to allow for more flexible generators (i.e. for different types of data) as well as user defined generators.

Currently, DRAWR has one generator implemented:
* SimpleGenerator

#### SimpleGenerator
This generator is designed to handle only the data above and left. It has 6 modes:
* 0 -- SimpleGeneratorModeStandard: Adds all data for the above and left pixel without doing any scaling.
* 1 -- SimpleGeneratorModeExclusive: When possible, it will use only pixels that exist in data for both the above and left pixels. If not possible, it defaults to left pixel data then above pixel data. Great for patterns.
* 2 -- SimpleGeneratorModeMultiplicative: Makes values which appear in data for both the upper and left pixels have a multiplicative bonus (about 100x bonus). 
* 3 -- SimpleGeneratorModeSuperMultiplicative: Makes values which appear in data for both the upper and left pixels have a multiplicative bonus (about 2000x bonus). 
* 4 -- SimpleGeneratorModeInverseMultiplicative: Makes values which appear in data for both the upper and left pixels have an inverse multiplicative bonus (about 1/5x bonus). Usually more chaotic.
* 5 -- SimpleGeneratorModeInverseSuperMultiplicative: Makes values which appear in data for both the upper and left pixels have an inverse multiplicative bonus (about 1/250x bonus). Usually more chaotic.

## License
DRAWR is licensed under an MIT license.
