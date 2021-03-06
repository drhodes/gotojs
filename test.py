#!/usr/bin/env python

import os, sys

TESTDIR = "./test/pass"
    
def testpath(path):
    # get the gotojs compiler output
    cmd = './gotojs -path="%s/%s"' % (TESTDIR, path)
    raw_result = os.popen(cmd).read()
    
    # write the output to the testfile.
    outname = "%s/%s/main.test" % (TESTDIR, path)
    of = open(outname, 'w')
    of.write(raw_result)
    of.close()
    
    # read the output from beautfier.
    cmd2 = "js-beautify %s" % (outname)
    beautified = os.popen(cmd2).read().strip()
    of = open(outname, 'w')
    of.write(beautified)
    of.close()

    # read the output from the control js file.
    control_path = "%s/%s/main.js" % (TESTDIR, path)
    control = open(control_path).read().strip()

    # compare the two outputs
    if control != beautified:    
        # if they are different, then fail.
        diff_cmd = "diff %s %s" % (outname, control_path)
        diff_output = os.popen(diff_cmd).read()
        print path
        print diff_output
    else:
        pass
        # else pass happy.
        # print "pass: " + path

def check_gotojs():
    if os.path.isfile("./gotojs"):
        return    
    print " gotojs not found, it's probably not built"
    print
    print " $ go build"
    sys.exit(1)
    
if __name__ == "__main__":
    check_gotojs()
    for d in os.listdir(TESTDIR):
        if d != "acid":
            testpath(d)
